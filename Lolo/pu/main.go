package main

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/gookit/slog"

	"gucooing/lolo/config"
	"gucooing/lolo/db"
	"gucooing/lolo/gateway"
	"gucooing/lolo/gdconf"
	"gucooing/lolo/logserver"
	"gucooing/lolo/pkg"
	"gucooing/lolo/pkg/log"
	"gucooing/lolo/sdk"
)

func main() {
	exit := func() {
		fmt.Printf("\n执行结束请输入任何键退出程序....")
		scanner := bufio.NewScanner(os.Stdin)
		for {
			scanner.Scan()
			return
		}
	}
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("\n程序异常退出,原因:")
			fmt.Println(err)
			exit()
		}
	}()
	newLolo()
}

func newLolo() {
	var filePath string
	var genConfig bool
	flag.StringVar(&filePath, "c", "./config.json", "配置文件路径")
	flag.BoolVar(&genConfig, "g", false, "是否生成默认配置文件")
	flag.Parse()

	if genConfig {
		fmt.Printf("生成默认配置文件\n")
		p, _ := json.MarshalIndent(config.DefaultConfig, "", "  ")
		cf, _ := os.Create(filePath)
		_, err := cf.Write(p)
		cf.Close()
		if err != nil {
			fmt.Printf("生成默认配置文件失败 %s \n请检查是否有权限\n", err.Error())
			return
		} else {
			fmt.Printf("生成默认配置文件成功 \n请修改后重新启动")
			return
		}
	}

	if err := config.LoadConfig(filePath); err != nil {
		panic(err)
	}
	log.NewApp()

	log.App.Info("Lolo Start")
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	// 初始化数据库
	log.App.Debug("开始初始化数据库")
	if err := db.NewDB(); err != nil {
		panic(fmt.Sprintf("初始化数据库失败:%s", err.Error()))
	} else {
		log.App.Debug("初始化数据库成功")
	}
	// 初始化gin
	ginRouter, httpServer := NewGin()
	// 初始化sdk
	s := sdk.New(ginRouter)
	// 初始化gateWay
	g := gateway.NewGateway(ginRouter)
	// 初始化logserver
	l := logserver.NewLogServer(ginRouter)
	// 初始化资源文件
	gdconf.LoadGameConfig()

	// 启动HTTP服务器
	log.App.Infof("ClientVersion:%s", pkg.ClientVersion)
	log.App.Infof("ServerVersion:%s", pkg.ServerVersion)
	log.App.Infof("Commit:%s", pkg.Commit)
	go func() {
		log.App.Info("Lolo Http Start!")
		if err := RunGin(httpServer); err != nil {
			if !errors.Is(http.ErrServerClosed, err) {
				log.App.Errorf("HTTP服务器错误:%s", err.Error())
				done <- syscall.SIGTERM
			}
		}
	}()

	// 启动GateWay服务器
	go func() {
		log.App.Info("Lolo GateWay Start!")
		if err := g.RunGateway(); err != nil {
			log.App.Errorf("GateWay服务器错误:%s", err.Error())
			done <- syscall.SIGTERM
		}
	}()

	// 启动LogServer服务器
	go func() {
		log.App.Info("Lolo LogServer Start!")
		if err := l.RunLogServer(); err != nil {
			log.App.Errorf("LogServer服务器错误:%s", err.Error())
			done <- syscall.SIGTERM
		}
	}()
	log.App.Info("Lolo Start!")

	// close
	clo := func() {
		_, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
		defer cancel()
		log.App.Info("Lolo Close")
		s.Close()
		g.Close()
		l.Close()
		log.Close()
		os.Exit(0)
	}

	go func() {
		select {
		case call := <-done:
			switch call {
			case syscall.SIGINT, syscall.SIGTERM:
				clo()
				return
			}
		}
	}()

	select {}
}

func NewGin() (*gin.Engine, *http.Server) {
	log.App.Debug("初始化gin服务")
	cfg := config.GetHttpNet()
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(
		gin.Recovery(),
		log.GinLog(log.App),
		gzip.Gzip(gzip.DefaultCompression))
	if log.App.Level >= slog.DebugLevel {
		pprof.Register(router)
	}
	addr := fmt.Sprintf("%s:%s", cfg.GetInnerIp(), cfg.GetInnerPort())
	log.App.Infof("监听地址: http://%s", addr)
	server := &http.Server{Addr: addr, Handler: router}
	log.App.Debug("gin服务初始化成功")
	return router, server
}

func RunGin(server *http.Server) error {
	log.App.Debug("启动http服务")
	return server.ListenAndServe()
}

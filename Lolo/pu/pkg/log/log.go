package log

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gookit/slog"
	"github.com/gookit/slog/handler"
	"github.com/gookit/slog/rotatefile"

	"gucooing/lolo/config"
)

var (
	App       *slog.SugaredLogger
	Gate      *slog.SugaredLogger
	Game      *slog.SugaredLogger
	ClientLog *slog.SugaredLogger
)

func Close() {
	if App != nil {
		App.Close()
	}
	if Gate != nil {
		Gate.Close()
	}
	if Game != nil {
		Game.Close()
	}
}

func NewApp() {
	conf := config.GetLog()
	App = slog.NewStdLogger(func(sl *slog.SugaredLogger) {
		f := sl.Formatter.(*slog.TextFormatter)
		f.EnableColor = true
		sl.ChannelName = conf.AppName
		sl.Level = conf.Level
	})
	addHandler(App, conf)
}

func NewGate() {
	conf := config.GetGateWay().GetLog()
	Gate = slog.NewStdLogger(func(sl *slog.SugaredLogger) {
		sl.ChannelName = conf.AppName
		sl.Level = conf.Level
	})
	addHandler(Gate, conf)
}

func NewGame() {
	conf := config.GetGame().GetLog()
	Game = slog.NewStdLogger(func(sl *slog.SugaredLogger) {
		sl.ChannelName = conf.AppName
		sl.Level = conf.Level
	})
	addHandler(Game, conf)
}

func NewClientLog() {
	conf := config.GetLogServer().GetLog()
	ClientLog = slog.NewStdLogger(func(sl *slog.SugaredLogger) {
		sl.ChannelName = conf.AppName
		sl.Level = conf.Level
	})
	addHandler(ClientLog, conf)
}

func addHandler(l *slog.SugaredLogger, conf *config.Log) {
	if conf.LogFile {
		l.AddHandler(handler.NewBuilder().
			WithLogfile(fmt.Sprintf("./log/%s.log", conf.AppName)).
			WithLogLevels(slog.AllLevels).
			WithBuffSize(1024 * 10).
			WithRotateTime(rotatefile.Every30Min).
			WithCompress(true).
			Build())
	}
}

func GinLog(l *slog.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()
		latency := time.Now().Sub(start)

		if raw != "" {
			path = path + "?" + raw
		}

		l.Debugf("HTTP [%s %s]|%3d||%13v|%15s|",
			c.Request.Method,
			path,
			c.Writer.Status(),
			latency,
			c.ClientIP())
	}
}

package gdconf

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/bytedance/sonic"

	"gucooing/lolo/config"
	"gucooing/lolo/pkg/log"
	"gucooing/lolo/protocol/excel"
)

var cc *GameConfig

type GameConfig struct {
	dataPath   string
	excelPath  string
	configPath string

	Excel  *Excel
	Config *Config

	Constant      *Constant
	ClientVersion *ClientVersion
}

func LoadGameConfig() *GameConfig {
	cfg := config.GetResources()
	c := new(GameConfig)
	c.dataPath = cfg.GetDataPath()
	c.excelPath = cfg.GetExcelPath()
	c.configPath = cfg.GetConfigPath()
	log.App.Debug("开始读取资源文件")
	startTime := time.Now()
	c.load()
	endTime := time.Now()
	cc = c
	runtime.GC()
	log.App.Debugf("读取资源完成,用时:%s", endTime.Sub(startTime))
	return c
}

func (g *GameConfig) load() {
	// 验证文件夹是否存在
	if dirInfo, err := os.Stat(g.dataPath); err != nil || !dirInfo.IsDir() {
		info := fmt.Sprintf("找不到文件夹:%s,err:%s", g.dataPath, err)
		panic(info)
	}
	g.dataPath += "/"

	// 验证文件夹是否存在
	if dirInfo, err := os.Stat(g.excelPath); err != nil || !dirInfo.IsDir() {
		info := fmt.Sprintf("找不到文件夹:%s,err:%s", g.excelPath, err)
		panic(info)
	}
	g.excelPath += "/"

	// 验证文件夹是否存在
	if dirInfo, err := os.Stat(g.configPath); err != nil || !dirInfo.IsDir() {
		info := fmt.Sprintf("找不到文件夹:%s,err:%s", g.configPath, err)
		panic(info)
	}
	g.configPath += "/"

	g.Excel = new(Excel)
	g.Config = new(Config)

	// data
	g.loadConstant()
	g.loadClientVersion()

	// 初始化excel
	g.loadHead()
	g.loadCharacter()
	g.loadItem()
	g.loadWeapon()

	// config
	g.loadSceneConfig()
}

type Excel struct {
	AllHeadDatas *excel.AllHeadDatas
	Character    *Character
	Item         *Item
	Weapon       *Weapon
}

type Config struct {
	SceneConfig *SceneConfig
}

func ReadJson[T any](path, name string, t *T) {
	file, err := os.ReadFile(path + name)
	if err != nil {
		log.App.Errorf("文件:%s 读取失败,err:%s", name, err)
		return
	}
	if err := sonic.Unmarshal(file, t); err != nil {
		log.App.Errorf("文件:%s 解析失败,err:%s", name, err)
		return
	}
	log.App.Infof("文件:%s 读取成功", name)
}

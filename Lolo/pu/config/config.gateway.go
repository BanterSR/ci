package config

import (
	"github.com/gookit/slog"
)

type GateWay struct {
	Log            *Log     `json:"Log"`
	OuterIp        string   `json:"OuterIp"`
	OuterPort      int      `json:"OuterPort"`
	OuterAddr      string   `json:"OuterAddr"`
	MaxPlayerNum   int64    `json:"MaxPlayerNum"`
	BlackCmd       []string `json:"BlackCmd"`
	IsLogMsgPlayer bool     `json:"IsLogMsgPlayer"`
}

var defaultGateWay = &GateWay{
	Log: &Log{
		Level:   slog.InfoLevel,
		LogFile: false,
		AppName: "Gate",
	},
	OuterIp:      "127.0.0.1",
	OuterPort:    11000,
	OuterAddr:    "0.0.0.0:11000",
	MaxPlayerNum: 0,
	BlackCmd: []string{
		"PlayerPingReq",
		"PlayerPingRsp",
		"PlayerSceneRecordReq",
		"PlayerSceneRecordRsp",
		"PlayerSceneSyncDataNotice",
	},
	IsLogMsgPlayer: false,
}

func GetGateWay() *GateWay {
	return GetConfig().GateWay
}

func (x *GateWay) GetLog() *Log {
	return x.Log
}

func (x *GateWay) GetOuterIp() string {
	return x.OuterIp
}

func (x *GateWay) GetOuterPort() int {
	return x.OuterPort
}

func (x *GateWay) GetOuterAddr() string {
	return x.OuterAddr
}

func (x *GateWay) GetMaxPlayerNum() int64 {
	return x.MaxPlayerNum
}

func (x *GateWay) GetBlackCmd() []string {
	return x.BlackCmd
}

func (x *GateWay) GetIsLogMsgPlayer() bool {
	return x.IsLogMsgPlayer
}

package config

import (
	"github.com/gookit/slog"
)

type LogServer struct {
	Log       *Log   `json:"Log"`
	OuterIp   string `json:"OuterIp"`
	OuterPort int    `json:"OuterPort"`
	OuterAddr string `json:"OuterAddr"`
}

var defaultLogServer = &LogServer{
	Log: &Log{
		Level:   slog.InfoLevel,
		LogFile: false,
		AppName: "Log",
	},
	OuterIp:   "127.0.0.1",
	OuterPort: 12000,
	OuterAddr: "0.0.0.0:12000",
}

func GetLogServer() *LogServer {
	if GetConfig().LogServer == nil {
		return defaultLogServer
	}
	return GetConfig().LogServer
}

func (x *LogServer) GetLog() *Log {
	return x.Log
}

func (x *LogServer) GetOuterIp() string {
	return x.OuterIp
}

func (x *LogServer) GetOuterPort() int {
	return x.OuterPort
}

func (x *LogServer) GetOuterAddr() string {
	return x.OuterAddr
}

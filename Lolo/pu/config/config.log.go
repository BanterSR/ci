package config

import (
	"github.com/gookit/slog"
)

type Log struct {
	Level   slog.Level `json:"Level"`
	LogFile bool       `json:"LogFile"`
	AppName string     `json:"AppName"`
}

var defaultLog = &Log{
	Level:   slog.InfoLevel,
	LogFile: false,
	AppName: "Lolo",
}

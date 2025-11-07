package model

import (
	"gucooing/lolo/gdconf"
	"gucooing/lolo/protocol/proto"
)

type BasicModel struct {
	PlayerName      string `json:"playerName,omitempty"` // 玩家昵称
	Level           uint32 `json:"level,omitempty"`      // 账号等级
	Exp             uint32 `json:"exp,omitempty"`        // 账号经验
	Head            uint32 `json:"head,omitempty"`       // 头像
	Sign            string `json:"sign,omitempty"`       // 签名
	Birthday        string `json:"birthday,omitempty"`   // 生日
	IsHideBirthday  bool   `json:"isHideBirthday,omitempty"`
	PhoneBackground uint32 `json:"phoneBackground,omitempty"` // 手机背景
}

func DefaultBasicModel() *BasicModel {
	return &BasicModel{
		PlayerName:      gdconf.GetConstant().DefaultPlayerName,
		Level:           gdconf.GetConstant().DefaultPlayerLevel,
		Exp:             gdconf.GetConstant().DefaultPlayerExp,
		Head:            gdconf.GetConstant().DefaultPlayerHead,
		Sign:            gdconf.GetConstant().DefaultPlayerSign,
		PhoneBackground: gdconf.GetConstant().DefaultPhoneBackground,
	}
}

func (s *Player) GetBasicModel() *BasicModel {
	if s == nil {
		return nil
	}
	if s.Basic == nil {
		s.Basic = DefaultBasicModel()
	}

	return s.Basic
}

func (s *Player) GetPbPlayerDropRateInfo() *proto.PlayerDropRateInfo {
	info := &proto.PlayerDropRateInfo{
		KillDropRate:     100,
		TreasureDropRate: 100,
	}
	return info
}

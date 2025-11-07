package model

import (
	"time"

	"gucooing/lolo/pkg/ofnet"
)

type Player struct {
	Conn          ofnet.Conn      `json:"-"`
	Online        bool            `json:"-"`
	NetFreeze     bool            `json:"-"`
	Created       time.Time       `json:"-"`                       // 创建时间
	Updated       time.Time       `json:"-"`                       // 更新时间
	UserId        uint32          `json:"-"`                       // 玩家id
	InstanceIndex uint32          `json:"instanceIndex,omitempty"` // 唯一索引生成
	Basic         *BasicModel     `json:"basic,omitempty"`         // 基础信息
	Item          *ItemModel      `json:"item,omitempty"`          // 背包
	Character     *CharacterModel `json:"character,omitempty"`     // 角色
	Team          *TeamModel      `json:"team,omitempty"`          // 队伍
}

func (s *Player) GetSeqId() uint32 {
	if s.Conn == nil {
		return 0
	}
	return s.Conn.GetSeqId()
}

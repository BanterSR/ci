package game

import (
	"gucooing/lolo/game/model"
	"gucooing/lolo/protocol/cmd"
	"gucooing/lolo/protocol/proto"
)

func (g *Game) ChatMsgRecordInitNotice(s *model.Player) {
	notice := &proto.ChatMsgRecordInitNotice{
		Status: proto.StatusCode_StatusCode_OK,
		Type:   0,
		Msg:    make([]*proto.ChatMsgData, 0),
	}
	defer g.send(s, cmd.ChatMsgRecordInitNotice, 0, notice)
}

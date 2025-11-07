package game

import (
	"gucooing/lolo/game/model"
	"gucooing/lolo/pkg/alg"
	"gucooing/lolo/protocol/cmd"
	"gucooing/lolo/protocol/proto"
)

func (g *Game) GetArchiveInfo(s *model.Player, msg *alg.GameMsg) {
	rsp := &proto.GetArchiveInfoRsp{
		Status: proto.StatusCode_StatusCode_OK,
		Key:    "",
		Value:  "",
	}
	g.send(s, cmd.GetArchiveInfoRsp, msg.PacketId, rsp)
}

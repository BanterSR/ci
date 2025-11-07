package game

import (
	"gucooing/lolo/game/model"
	"gucooing/lolo/pkg/alg"
	"gucooing/lolo/protocol/cmd"
	"gucooing/lolo/protocol/proto"
)

func (g *Game) GachaList(s *model.Player, msg *alg.GameMsg) {
	rsp := &proto.GachaListRsp{
		Status: proto.StatusCode_StatusCode_OK,
		Gachas: make([]*proto.GachaInfo, 0),
	}
	defer g.send(s, cmd.GachaListRsp, msg.PacketId, rsp)
}

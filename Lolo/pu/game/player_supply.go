package game

import (
	"gucooing/lolo/game/model"
	"gucooing/lolo/pkg/alg"
	"gucooing/lolo/protocol/cmd"
	"gucooing/lolo/protocol/proto"
)

func (g *Game) SupplyBoxInfo(s *model.Player, msg *alg.GameMsg) {
	rsp := &proto.SupplyBoxInfoRsp{
		Status:         proto.StatusCode_StatusCode_OK,
		NextRewardTime: 0,
	}
	defer g.send(s, cmd.SupplyBoxInfoRsp, msg.PacketId, rsp)
}

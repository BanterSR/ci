package game

import (
	"gucooing/lolo/game/model"
	"gucooing/lolo/pkg/alg"
	"gucooing/lolo/protocol/cmd"
	"gucooing/lolo/protocol/proto"
)

func (g *Game) ShopInfo(s *model.Player, msg *alg.GameMsg) {
	req := msg.Body.(*proto.ShopInfoReq)
	rsp := &proto.ShopInfoRsp{
		Status: proto.StatusCode_StatusCode_OK,
		ShopId: req.ShopId,
		Grids:  make([]*proto.ShopGrid, 0),
	}
	defer g.send(s, cmd.ShopInfoRsp, msg.PacketId, rsp)
}

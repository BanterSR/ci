package game

import (
	"gucooing/lolo/game/model"
	"gucooing/lolo/pkg/alg"
	"gucooing/lolo/protocol/cmd"
	"gucooing/lolo/protocol/proto"
)

func (g *Game) GetAllCharacterEquip(s *model.Player, msg *alg.GameMsg) {
	rsp := &proto.GetAllCharacterEquipRsp{
		Status: proto.StatusCode_StatusCode_OK,
		Items:  make([]*proto.ItemDetail, 0),
	}
	defer g.send(s, cmd.GetAllCharacterEquipRsp, msg.PacketId, rsp)
}

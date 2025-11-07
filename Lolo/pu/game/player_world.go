package game

import (
	"gucooing/lolo/game/model"
	"gucooing/lolo/pkg/alg"
	"gucooing/lolo/protocol/cmd"
	"gucooing/lolo/protocol/proto"
)

func (g *Game) WorldLevelAchieveList(s *model.Player, msg *alg.GameMsg) {
	rsp := &proto.WorldLevelAchieveListRsp{
		Status:            proto.StatusCode_StatusCode_OK,
		Achieves:          make([]*proto.Achieve, 0),
		UnlockWorldLevels: make([]uint32, 0),
	}
	defer g.send(s, cmd.WorldLevelAchieveListRsp, msg.PacketId, rsp)
}

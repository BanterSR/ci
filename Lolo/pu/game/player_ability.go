package game

import (
	"gucooing/lolo/game/model"
	"gucooing/lolo/pkg/alg"
	"gucooing/lolo/protocol/cmd"
	"gucooing/lolo/protocol/proto"
)

func (g *Game) AbilityBadgeList(s *model.Player, msg *alg.GameMsg) {
	rsp := &proto.AbilityBadgeListRsp{
		Status:                       proto.StatusCode_StatusCode_OK,
		AbilityBadgePages:            make([]*proto.AbilityBadgePageInfo, 0),
		AbilityBadgeAchieves:         make([]*proto.Achieve, 0),
		AbilityBadgeRewardAchieveIds: make([]uint32, 0),
	}
	defer g.send(s, cmd.AbilityBadgeListRsp, msg.PacketId, rsp)
}

func (g *Game) PlayerAbilityList(s *model.Player, msg *alg.GameMsg) {
	rsp := &proto.PlayerAbilityListRsp{
		Status:      proto.StatusCode_StatusCode_OK,
		AbilityList: make([]*proto.PlayerAbility, 0),
	}
	defer g.send(s, cmd.PlayerAbilityListRsp, msg.PacketId, rsp)
}

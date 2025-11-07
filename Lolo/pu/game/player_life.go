package game

import (
	"gucooing/lolo/game/model"
	"gucooing/lolo/pkg/alg"
	"gucooing/lolo/protocol/cmd"
	"gucooing/lolo/protocol/proto"
)

func (g *Game) GetLifeInfo(s *model.Player, msg *alg.GameMsg) {
	req := msg.Body.(*proto.GetLifeInfoReq)
	rsp := &proto.GetLifeInfoRsp{
		Status: proto.StatusCode_StatusCode_OK,
		LifeBaseInfo: &proto.LifeBaseInfo{
			LifeType:         req.GetLifeType(),
			ProficiencyValue: 0,
			Level:            1,
		},
		LifeSkill: &proto.LifeSkill{
			LifeType:      req.GetLifeType(),
			LifeSkillInfo: nil,
		},
		LifeAchieve: &proto.LifeAchieve{
			LifeType:                 req.GetLifeType(),
			LifeAchieve_:             nil,
			RewardedLifeAchieveIdLst: nil,
		},
	}
	defer g.send(s, cmd.GetLifeInfoRsp, msg.PacketId, rsp)
}

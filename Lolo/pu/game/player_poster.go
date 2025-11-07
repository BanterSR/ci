package game

import (
	"gucooing/lolo/game/model"
	"gucooing/lolo/pkg/alg"
	"gucooing/lolo/protocol/cmd"
	"gucooing/lolo/protocol/proto"
)

func (g *Game) PosterIllustrationList(s *model.Player, msg *alg.GameMsg) {
	rsp := &proto.PosterIllustrationListRsp{
		Status:              proto.StatusCode_StatusCode_OK,
		PosterIllustrations: make([]*proto.PosterIllustration, 0),
	}
	defer g.send(s, cmd.PosterIllustrationListRsp, msg.PacketId, rsp)
}

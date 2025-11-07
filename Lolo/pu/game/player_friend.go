package game

import (
	"gucooing/lolo/game/model"
	"gucooing/lolo/pkg/alg"
	"gucooing/lolo/protocol/cmd"
	"gucooing/lolo/protocol/proto"
)

func (g *Game) Friend(s *model.Player, msg *alg.GameMsg) {
	// req := msg.Body.(*proto.FriendReq)
	rsp := &proto.FriendRsp{
		Status: proto.StatusCode_StatusCode_OK,
		Info:   make([]*proto.FriendBriefInfo, 0),
	}
	defer g.send(s, cmd.FriendRsp, msg.PacketId, rsp)
}

func (g *Game) WishListByFriendId(s *model.Player, msg *alg.GameMsg) {
	// req := msg.Body.(*proto.FriendReq)
	rsp := &proto.WishListByFriendIdRsp{
		Status:        proto.StatusCode_StatusCode_OK,
		PlayerId:      0,
		WishList:      make([]*proto.WishListInfo, 0),
		WeekSendCount: 0,
	}
	defer g.send(s, cmd.WishListByFriendIdRsp, msg.PacketId, rsp)
}

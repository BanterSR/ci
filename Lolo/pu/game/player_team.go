package game

import (
	"gucooing/lolo/game/model"
	"gucooing/lolo/pkg/alg"
	"gucooing/lolo/pkg/log"
	"gucooing/lolo/protocol/cmd"
	"gucooing/lolo/protocol/proto"
)

func (g *Game) UpdateTeam(s *model.Player, msg *alg.GameMsg) {
	req := msg.Body.(*proto.UpdateTeamReq)
	rsp := &proto.UpdateTeamRsp{
		Status: proto.StatusCode_StatusCode_OK,
	}
	defer g.send(s, cmd.UpdateTeamRsp, msg.PacketId, rsp)
	// 更新队伍
	upChar := func(target *uint32, char uint32) bool {
		*target = char
		return true
	}
	teamInfo := s.GetTeamInfo()
	upChar(&teamInfo.Char1, req.Char1)
	upChar(&teamInfo.Char2, req.Char2)
	upChar(&teamInfo.Char3, req.Char3)

	scenePlayer := g.getWordInfo().getScenePlayer(s)
	if scenePlayer == nil ||
		scenePlayer.channelInfo == nil {
		rsp.Status = proto.StatusCode_StatusCode_PLAYER_NOT_IN_CHANNEL
		log.Game.Warnf("玩家:%v没有加入房间", s.UserId)
		return
	}
	scenePlayer.channelInfo.serverSceneSyncChan <- &ServerSceneSyncCtx{
		ScenePlayer: scenePlayer,
		ActionType:  proto.SceneActionType_SceneActionType_UPDATE_TEAM,
	}
}

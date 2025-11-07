package game

import (
	"gucooing/lolo/game/model"
	"gucooing/lolo/pkg/alg"
	"gucooing/lolo/pkg/log"
	"gucooing/lolo/protocol/cmd"
	"gucooing/lolo/protocol/proto"
)

func (g *Game) PlayerSceneRecord(s *model.Player, msg *alg.GameMsg) {
	req := msg.Body.(*proto.PlayerSceneRecordReq)
	rsp := &proto.PlayerSceneRecordRsp{
		Status: proto.StatusCode_StatusCode_OK,
	}
	defer g.send(s, cmd.PlayerSceneRecordRsp, msg.PacketId, rsp)
	scenePlayer := g.getWordInfo().getScenePlayer(s)
	if scenePlayer == nil ||
		scenePlayer.channelInfo == nil {
		rsp.Status = proto.StatusCode_StatusCode_PLAYER_NOT_IN_CHANNEL
		log.Game.Warnf("玩家:%v没有加入房间", s.UserId)
		return
	}
	scenePlayer.channelInfo.addSceneSyncDataChan <- &proto.SceneSyncData{
		PlayerId: s.UserId,
		Data:     []*proto.PlayerRecorderData{req.Data},
	}
}

func (g *Game) SceneProcessList(s *model.Player, msg *alg.GameMsg) {
	rsp := &proto.SceneProcessListRsp{
		Status:           proto.StatusCode_StatusCode_OK,
		SceneProcessList: make([]*proto.SceneProcess, 0),
	}
	defer g.send(s, cmd.SceneProcessListRsp, msg.PacketId, rsp)
	rsp.SceneProcessList = append(rsp.SceneProcessList, &proto.SceneProcess{
		SceneId: 9999,
		Process: 1,
	})
	rsp.SceneProcessList = append(rsp.SceneProcessList, &proto.SceneProcess{
		SceneId: 1,
		Process: 1,
	})
	rsp.SceneProcessList = append(rsp.SceneProcessList, &proto.SceneProcess{
		SceneId: 100,
		Process: 1,
	})
}

func (g *Game) SendAction(s *model.Player, msg *alg.GameMsg) {
	req := msg.Body.(*proto.SendActionReq)
	rsp := &proto.SendActionRsp{
		Status: proto.StatusCode_StatusCode_OK,
	}
	defer g.send(s, cmd.SendActionRsp, msg.PacketId, rsp)
	scenePlayer := g.getWordInfo().getScenePlayer(s)
	if scenePlayer == nil ||
		scenePlayer.channelInfo == nil {
		rsp.Status = proto.StatusCode_StatusCode_PLAYER_NOT_IN_CHANNEL
		log.Game.Warnf("玩家:%v没有加入房间", s.UserId)
		return
	}
	scenePlayer.channelInfo.actionSyncChan <- &ActionSyncCtx{
		ScenePlayer: scenePlayer,
		ActionId:    req.ActionId,
	}
}

func (g *Game) GenericSceneB(s *model.Player, msg *alg.GameMsg) {
	req := msg.Body.(*proto.GenericSceneBReq)
	rsp := &proto.GenericSceneBRsp{
		Status:       proto.StatusCode_StatusCode_OK,
		GenericMsgId: req.GenericMsgId,
		Params:       make([]*proto.CommonParam, 0),
	}
	defer g.send(s, cmd.GenericSceneBRsp, msg.PacketId, rsp)
}

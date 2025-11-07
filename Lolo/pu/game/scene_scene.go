package game

import (
	"errors"

	"gucooing/lolo/game/model"
	"gucooing/lolo/gdconf"
	"gucooing/lolo/pkg/log"
	"gucooing/lolo/protocol/config"
	"gucooing/lolo/protocol/proto"
)

var (
	minChannelId = uint32(1)
	maxChannelId = uint32(99999)
)

type WordInfo struct {
	allScene       map[uint32]*SceneInfo   // 整个服务上的全部场景
	allScenePlayer map[uint32]*ScenePlayer // 整个服务上的全部场景玩家对象
}

type SceneInfo struct {
	cfg        *config.SceneInfo       // 场景配置
	SceneId    uint32                  // 场景id
	allChannel map[uint32]*ChannelInfo // 全部房间
}

// 场景中玩家对象

type ScenePlayer struct {
	*model.Player
	Pos         *proto.Vector3
	Rot         *proto.Vector3
	SceneId     uint32
	ChannelId   uint32
	channelInfo *ChannelInfo // 绑定的房间
}

func (g *Game) getWordInfo() *WordInfo {
	if g.wordInfo == nil {
		g.wordInfo = new(WordInfo)
	}
	return g.wordInfo
}

func (w *WordInfo) getAllSceneInfo() map[uint32]*SceneInfo {
	if w.allScene == nil {
		w.allScene = make(map[uint32]*SceneInfo)
	}
	return w.allScene
}

func (w *WordInfo) getSceneInfo(sceneId uint32) *SceneInfo {
	list := w.getAllSceneInfo()
	if info, ok := list[sceneId]; ok {
		return info
	}
	cfg := gdconf.GetSceneInfo(sceneId)
	if cfg == nil {
		return nil
	}
	info := &SceneInfo{
		cfg:        cfg,
		SceneId:    sceneId,
		allChannel: make(map[uint32]*ChannelInfo),
	}
	list[sceneId] = info
	return info
}

func (w *WordInfo) getAllScenePlayer() map[uint32]*ScenePlayer {
	if w.allScenePlayer == nil {
		w.allScenePlayer = make(map[uint32]*ScenePlayer)
	}
	return w.allScenePlayer
}

func (w *WordInfo) getScenePlayer(player *model.Player) *ScenePlayer {
	list := w.getAllScenePlayer()
	if info, ok := list[player.UserId]; ok {
		return info
	}
	// 场景中没有该玩家
	return nil
}

// 获取目标场景下的全部房间
func (s *SceneInfo) getAllSceneChannel() map[uint32]*ChannelInfo {
	if s.allChannel == nil {
		s.allChannel = make(map[uint32]*ChannelInfo)
	}
	return s.allChannel
}

func (s *SceneInfo) getSceneChannel(channelId uint32) (*ChannelInfo, error) {
	list := s.getAllSceneChannel()
	if info, ok := list[channelId]; ok {
		return info, nil
	}
	// add SceneChannel
	if channelId >= minChannelId && channelId <= maxChannelId {
		info := s.newChannelInfo(channelId)
		list[channelId] = info
		return info, nil
	}
	if channelId >= 1000000 {
		return nil, errors.New("暂未实现私人房间")
	}
	return nil, errors.New("没有该房间")
}

// 添加场景玩家对象
func (w *WordInfo) addScenePlayer(player *model.Player) *ScenePlayer {
	list := w.getAllScenePlayer()
	if info, ok := list[player.UserId]; ok {
		return info
	}
	defaultSceneId := gdconf.GetConstant().DefaultSceneId
	defaultChannelId := gdconf.GetConstant().DefaultChannelId

	sceneInfo := w.getSceneInfo(defaultSceneId)
	if sceneInfo == nil {
		log.Game.Error("默认场景不存在！请检查默认场景配置是否正确")
		return nil
	}
	_, err := sceneInfo.getSceneChannel(defaultChannelId)
	if err != nil {
		log.Game.Error(err.Error())
		return nil
	}
	pos, rot := gdconf.GetSceneInfoRandomBorn(sceneInfo.cfg)

	info := &ScenePlayer{
		Player:    player,
		Pos:       gdconf.ConfigVector3ToProtoVector3(pos),
		Rot:       gdconf.ConfigVector4ToProtoVector3(rot),
		SceneId:   defaultSceneId,   // 默认场景
		ChannelId: defaultChannelId, // 默认房间
	}
	list[player.UserId] = info
	return info
}

// 加入房间
func (g *Game) joinSceneChannel(s *model.Player) {
	scenePlayer := g.getWordInfo().getScenePlayer(s)
	if scenePlayer == nil {
		log.Game.Warnf("玩家:%v没有准备好加入房间", s.UserId)
		return
	}
	sceneInfo := g.getWordInfo().getSceneInfo(scenePlayer.SceneId)
	if sceneInfo == nil {
		log.Game.Errorf("场景:%v不存在！", scenePlayer.SceneId)
		return
	}
	channelInfo, err := sceneInfo.getSceneChannel(scenePlayer.ChannelId)
	if err != nil {
		log.Game.Errorf("场景:%v,房间:%v不存在！", scenePlayer.SceneId, scenePlayer.ChannelId)
		return
	}
	scenePlayer.channelInfo = channelInfo
	channelInfo.addScenePlayerChan <- scenePlayer
}

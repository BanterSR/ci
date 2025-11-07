package gdconf

import (
	"math/rand"

	"gucooing/lolo/protocol/config"
	"gucooing/lolo/protocol/proto"
)

type SceneConfig struct {
	SceneConfig    *config.SceneConfig
	SceneBySceneId map[int32]*config.SceneInfo
}

func (g *GameConfig) loadSceneConfig() {
	info := &SceneConfig{
		SceneConfig:    new(config.SceneConfig),
		SceneBySceneId: make(map[int32]*config.SceneInfo),
	}
	g.Config.SceneConfig = info
	name := "ScenesConfigAsset.json"
	ReadJson(g.configPath, name, &info.SceneConfig)
	for _, v := range info.SceneConfig.GetScenes() {
		info.SceneBySceneId[v.ID] = v
	}
}

func GetSceneInfo(sceneId uint32) *config.SceneInfo {
	return cc.Config.SceneConfig.SceneBySceneId[int32(sceneId)]
}

func GetSceneInfoRandomBorn(info *config.SceneInfo) (*config.Vector3, *config.Vector4) {
	n := len(info.GetBorn())
	if n == 0 {
		return nil, nil
	}
	bornInfo := info.GetBorn()[rand.Intn(n)]
	return bornInfo.Position, bornInfo.Rotation
}

func ConfigVector3ToProtoVector3(s *config.Vector3) *proto.Vector3 {
	return &proto.Vector3{
		X:             int32(s.GetX() * 100),
		Y:             int32(s.GetY() * 100),
		Z:             int32(s.GetZ() * 100),
		DecimalPlaces: 0,
	}
}

func ConfigVector4ToProtoVector3(s *config.Vector4) *proto.Vector3 {
	return &proto.Vector3{
		X:             int32(s.GetX() * 100),
		Y:             int32(s.GetY() * 100),
		Z:             int32(s.GetZ() * 100),
		DecimalPlaces: 0,
	}
}

package gdconf

type Constant struct {
	DefaultCharacter       []uint32 `json:"DefaultCharacter"`
	DefaultBadge           uint32   `json:"DefaultBadge"`
	DefaultUmbrellaId      uint32   `json:"DefaultUmbrellaId"`
	EquipmentPresetNum     int      `json:"EquipmentPresetNum"`
	OutfitPresetNum        int      `json:"OutfitPresetNum"`
	DefaultPlayerName      string   `json:"DefaultPlayerName"`
	DefaultPlayerLevel     uint32   `json:"DefaultPlayerLevel"`
	DefaultPlayerExp       uint32   `json:"DefaultPlayerExp"`
	DefaultPlayerSign      string   `json:"DefaultPlayerSign"`
	DefaultPlayerHead      uint32   `json:"DefaultPlayerHead"`
	DefaultInstanceIndex   uint32   `json:"DefaultInstanceIndex"`
	DefaultPhoneBackground uint32   `json:"DefaultPhoneBackground"`
	DefaultSceneId         uint32   `json:"DefaultSceneId"`
	DefaultChannelId       uint32   `json:"DefaultChannelId"`
	ChannelTick            int      `json:"ChannelTick"`
}

func (g *GameConfig) loadConstant() {
	g.Constant = new(Constant)
	ReadJson(g.dataPath, "Constant.json", &g.Constant)
}

func GetConstant() *Constant {
	return cc.Constant
}

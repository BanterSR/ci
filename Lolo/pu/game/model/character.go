package model

import (
	"gucooing/lolo/gdconf"
	"gucooing/lolo/pkg/alg"
	"gucooing/lolo/pkg/log"
	"gucooing/lolo/protocol/proto"
)

type CharacterModel struct {
	CharacterMap map[uint32]*CharacterInfo `json:"characterMap,omitempty"`
}

func DefaultCharacterModel() *CharacterModel {
	info := &CharacterModel{
		CharacterMap: make(map[uint32]*CharacterInfo),
	}
	for _, characterId := range gdconf.GetConstant().DefaultCharacter {
		characterInfo := NewCharacterInfo(characterId)
		if characterInfo == nil {
			log.Game.Errorf("初始化默认角色:%v失败", characterId)
			continue
		}
		info.CharacterMap[characterId] = characterInfo
	}
	return info
}

func (s *Player) GetCharacterModel() *CharacterModel {
	if s == nil {
		return nil
	}
	if s.Character == nil {
		s.Character = DefaultCharacterModel()
	}
	return s.Character
}

type CharacterInfo struct {
	CharacterId               uint32                      `json:"characterId,omitempty"`               // 角色id
	Level                     uint32                      `json:"level,omitempty"`                     // 角色等级
	Exp                       uint32                      `json:"exp,omitempty"`                       // 角色经验
	Star                      uint32                      `json:"star,omitempty"`                      // 角色星级
	CharacterAppearance       *CharacterAppearance        `json:"characterAppearance,omitempty"`       // 角色外貌
	CharacterSkillList        map[uint32]*CharacterSkill  `json:"characterSkillList,omitempty"`        // 角色技能
	InUseEquipmentPresetIndex uint32                      `json:"inUseEquipmentPresetIndex,omitempty"` // 当前装备表
	EquipmentPresetList       map[uint32]*EquipmentPreset `json:"equipmentPresets,omitempty"`          // 角色装备预设表
	InUseOutfitPresetIndex    uint32                      `json:"inUseOutfitPresetIndex,omitempty"`    // 当前服装表
	OutfitPresetList          map[uint32]*OutfitPreset    `json:"outfitPresets,omitempty"`             // 角色服装预设表
}

func NewCharacterInfo(characterId uint32) *CharacterInfo {
	conf := gdconf.GetCharacter(int32(characterId))
	if conf == nil {
		return nil
	}
	info := &CharacterInfo{
		CharacterId: uint32(conf.ID),
		Level:       1,
		Exp:         0,
		Star:        0,
	}

	return info
}

func (s *Player) GetCharacterMap() map[uint32]*CharacterInfo {
	info := s.GetCharacterModel()
	if info == nil {
		return nil
	}
	if info.CharacterMap == nil {
		log.Game.Errorf("玩家:%v没有默认角色", s.UserId)
		return nil
	}
	return info.CharacterMap
}

func (s *Player) GetCharacterInfo(characterId uint32) *CharacterInfo {
	list := s.GetCharacterMap()
	if list == nil {
		return nil
	}
	return list[characterId]
}

func (s *Player) AddCharacter(characterId uint32) bool {
	list := s.GetCharacterMap()
	if list == nil {
		return false
	}
	conf := gdconf.GetCharacter(int32(characterId))
	if conf == nil {
		log.Game.Warnf("尝试添加不存在的角色:%v", characterId)
		return false
	}
	if list[characterId] != nil {
		return true
	}
	list[characterId] = NewCharacterInfo(characterId)
	return true
}

func (s *Player) GetAllPbCharacter() []*proto.Character {
	list := make([]*proto.Character, 0)
	for _, characterInfo := range s.GetCharacterMap() {
		alg.AddList(&list, s.GetPbCharacter(characterInfo))
	}
	return list
}

func (s *Player) GetPbCharacter(info *CharacterInfo) *proto.Character {
	if info == nil {
		return nil
	}
	pbInfo := &proto.Character{
		CharacterId:               info.CharacterId,
		Level:                     info.Level,
		MaxLevel:                  0,
		Exp:                       info.Exp,
		Star:                      info.Star,
		EquipmentPresets:          s.GetPbEquipmentPresets(info),
		InUseEquipmentPresetIndex: info.InUseEquipmentPresetIndex,
		OutfitPresets:             s.GetPbOutfitPresets(info),
		InUseOutfitPresetIndex:    info.InUseOutfitPresetIndex,
		GatherWeapon:              0,
		CharacterAppearance:       s.GetPbCharacterAppearance(info),
		CharacterSkillList:        s.GetPbCharacterSkillList(info),
		RewardedAchievementIdLst:  nil,
		IsUnlockPayment:           false,
		RewardIndexLst:            nil,
		MpGameWeapon:              0,
	}

	return pbInfo
}

type CharacterAppearance struct {
	Badge                      uint32 `json:"badge,omitempty"`      // 徽章
	UmbrellaId                 uint32 `json:"umbrellaId,omitempty"` // 伞？
	InsectNetInstanceId        uint32 `json:"insectNetInstanceId,omitempty"`
	LoggingAxeInstanceId       uint32 `json:"loggingAxeInstanceId,omitempty"`
	WaterBottleInstanceId      uint32 `json:"waterBottleInstanceId,omitempty"`
	MiningHammerInstanceId     uint32 `json:"miningHammerInstanceId,omitempty"`
	CollectionGlovesInstanceId uint32 `json:"collectionGlovesInstanceId,omitempty"`
	FishingRodInstanceId       uint32 `json:"fishingRodInstanceId,omitempty"`
}

func (s *Player) NewCharacterAppearance(characterId uint32) *CharacterAppearance {
	info := &CharacterAppearance{
		Badge:                      gdconf.GetConstant().DefaultBadge,
		UmbrellaId:                 gdconf.GetConstant().DefaultUmbrellaId,
		InsectNetInstanceId:        0,
		LoggingAxeInstanceId:       0,
		WaterBottleInstanceId:      0,
		MiningHammerInstanceId:     0,
		CollectionGlovesInstanceId: 0,
		FishingRodInstanceId:       0,
	}
	s.GetItemModel().AddItemBadge(info.Badge, 1)
	s.GetItemModel().AddItemUmbrella(info.UmbrellaId, 1)

	return info
}

func (s *Player) GetCharacterAppearance(info *CharacterInfo) *CharacterAppearance {
	if info.CharacterAppearance == nil {
		info.CharacterAppearance = s.NewCharacterAppearance(info.CharacterId)
	}
	return info.CharacterAppearance
}

func (s *Player) GetPbCharacterAppearance(characterInfo *CharacterInfo) *proto.CharacterAppearance {
	info := s.GetCharacterAppearance(characterInfo)
	if info == nil {
		return nil
	}
	pbInfo := &proto.CharacterAppearance{
		Badge:                      info.Badge,
		UmbrellaId:                 info.UmbrellaId,
		InsectNetInstanceId:        info.InsectNetInstanceId,
		LoggingAxeInstanceId:       info.LoggingAxeInstanceId,
		WaterBottleInstanceId:      info.WaterBottleInstanceId,
		MiningHammerInstanceId:     info.MiningHammerInstanceId,
		CollectionGlovesInstanceId: info.CollectionGlovesInstanceId,
		FishingRodInstanceId:       info.FishingRodInstanceId,
	}

	return pbInfo
}

type CharacterSkill struct {
	SkillId    uint32 `json:"skillId,omitempty"`
	SkillLevel uint32 `json:"skillLevel,omitempty"`
}

func NewCharacterSkillList(id uint32) map[uint32]*CharacterSkill {
	conf := gdconf.GetCharacter(int32(id))
	if conf == nil {
		log.Game.Warnf("添加技能的角色不存在id:%v", id)
		return nil
	}
	list := make(map[uint32]*CharacterSkill)
	addSkill := func(skillId uint32) {
		list[skillId] = &CharacterSkill{
			SkillId:    skillId,
			SkillLevel: 1,
		}
	}
	for _, skillId := range conf.GetSpellIDs() {
		addSkill(uint32(skillId))
	}
	for _, skillId := range conf.GetExSpellIDs() {
		addSkill(uint32(skillId))
	}
	return list
}

func (s *Player) GetCharacterSkillList(info *CharacterInfo) map[uint32]*CharacterSkill {
	if info.CharacterSkillList == nil {
		info.CharacterSkillList = NewCharacterSkillList(info.CharacterId)
	}
	return info.CharacterSkillList
}

func (s *Player) GetPbCharacterSkillList(characterInfo *CharacterInfo) []*proto.CharacterSkill {
	infoList := s.GetCharacterSkillList(characterInfo)
	if infoList == nil {
		return nil
	}
	pbInfoList := make([]*proto.CharacterSkill, 0)
	for _, info := range infoList {
		alg.AddList(&pbInfoList, &proto.CharacterSkill{
			SkillId:    info.SkillId,
			SkillLevel: info.SkillLevel,
		})
	}

	return pbInfoList
}

type EquipmentPreset struct {
	PresetIndex uint32                                       `json:"presetIndex,omitempty"`
	Weapon      uint32                                       `json:"weapon,omitempty"`
	Armors      map[proto.EEquipType]*ArmorInfo              `json:"armors,omitempty"`
	Posters     map[proto.PosterInfo_PosterIndex]*PosterInfo `json:"posters,omitempty"`
}

type ArmorInfo struct {
	EquipType proto.EEquipType `json:"equipType,omitempty"`
	ArmorId   uint32           `json:"armorId,omitempty"`
}

func (a *ArmorInfo) ArmorInfo() *proto.ArmorInfo {
	return &proto.ArmorInfo{
		EquipType: a.EquipType,
		ArmorId:   a.ArmorId,
	}
}

type PosterInfo struct {
	PosterIndex proto.PosterInfo_PosterIndex `json:"posterIndex,omitempty"`
	PosterId    uint32                       `json:"posterId,omitempty"`
}

func (a *PosterInfo) PosterInfo() *proto.PosterInfo {
	return &proto.PosterInfo{
		PosterIndex: a.PosterIndex,
		PosterId:    a.PosterId,
	}
}

func (s *Player) NewEquipmentPresetList(id uint32) map[uint32]*EquipmentPreset {
	conf := gdconf.GetCharacter(int32(id))
	if conf == nil {
		log.Game.Warnf("角色:%v获取初始装备套装失败", id)
		return nil
	}
	list := make(map[uint32]*EquipmentPreset)
	for i := 0; i < gdconf.GetConstant().EquipmentPresetNum; i++ {
		info := &EquipmentPreset{
			PresetIndex: uint32(i),
			Weapon:      0,
			Armors:      make(map[proto.EEquipType]*ArmorInfo),
			Posters:     make(map[proto.PosterInfo_PosterIndex]*PosterInfo),
		}
		if i == 0 {
			// 添加武器
			itemWeapon := s.GetItemModel().AddItemWeaponInfo(uint32(conf.DefaultWeaponID))
			if itemWeapon == nil {
				log.Game.Warnf("角色:%v,添加默认武器:%v失败", id, conf.DefaultWeaponID)
				return nil
			}
			itemWeapon.SetWearerId(id)
			info.Weapon = itemWeapon.InstanceId
		}
		// 添加盔甲
		for _, tag := range proto.EEquipType_value {
			info.Armors[proto.EEquipType(tag)] = &ArmorInfo{
				EquipType: proto.EEquipType(tag),
				ArmorId:   0,
			}
		}
		// 添加海报
		for _, index := range proto.PosterInfo_PosterIndex_value {
			info.Posters[proto.PosterInfo_PosterIndex(index)] = &PosterInfo{
				PosterIndex: proto.PosterInfo_PosterIndex(index),
				PosterId:    0,
			}
		}
		list[uint32(i)] = info
	}

	return list
}

func (s *Player) GetEquipmentPresetList(info *CharacterInfo) map[uint32]*EquipmentPreset {
	if info.EquipmentPresetList == nil {
		info.EquipmentPresetList = s.NewEquipmentPresetList(info.CharacterId)
	}
	return info.EquipmentPresetList
}

func (s *Player) GetEquipmentPreset(info *CharacterInfo, index uint32) *EquipmentPreset {
	list := s.GetEquipmentPresetList(info)
	return list[index]
}

func (s *Player) GetPbEquipmentPresets(characterInfo *CharacterInfo) []*proto.EquipmentPreset {
	infoList := s.GetEquipmentPresetList(characterInfo)
	if infoList == nil {
		return nil
	}
	pbInfoList := make([]*proto.EquipmentPreset, 0)
	for _, info := range infoList {
		pbInfo := &proto.EquipmentPreset{
			PresetIndex: info.PresetIndex,
			Weapon:      info.Weapon,
			Armors:      make([]*proto.ArmorInfo, 0),
			Posters:     make([]*proto.PosterInfo, 0),
		}
		for _, armor := range info.Armors {
			alg.AddList(&pbInfo.Armors, armor.ArmorInfo())
		}
		for _, poster := range info.Posters {
			alg.AddList(&pbInfo.Posters, poster.PosterInfo())
		}
		alg.AddList(&pbInfoList, pbInfo)
	}

	return pbInfoList
}

type OutfitPreset struct {
	PresetIndex            uint32          `json:"presetIndex,omitempty"`
	Hat                    uint32          `json:"hat,omitempty"`
	Hair                   uint32          `json:"hair,omitempty"`
	Clothes                uint32          `json:"clothes,omitempty"`
	Ornament               uint32          `json:"ornament,omitempty"`
	HatDyeSchemeIndex      uint32          `json:"hatDyeSchemeIndex,omitempty"`
	HairDyeSchemeIndex     uint32          `json:"hairDyeSchemeIndex,omitempty"`
	ClothesDyeSchemeIndex  uint32          `json:"clothesDyeSchemeIndex,omitempty"`
	OrnamentDyeSchemeIndex uint32          `json:"ornamentDyeSchemeIndex,omitempty"`
	OutfitHideInfo         *OutfitHideInfo `json:"outfitHideInfo,omitempty"`
}

type OutfitHideInfo struct {
	HideOrn   bool `json:"hideOrn,omitempty"`
	HideBraid bool `json:"hideBraid,omitempty"`
}

func (o *OutfitHideInfo) OutfitHideInfo() *proto.OutfitHideInfo {
	return &proto.OutfitHideInfo{
		HideOrn:   o.HideOrn,
		HideBraid: o.HideBraid,
	}
}

func (s *Player) NewOutfitPresetList(id uint32) map[uint32]*OutfitPreset {
	conf := gdconf.GetCharacter(int32(id))
	if conf == nil {
		log.Game.Warnf("角色:%v获取初始服装套装失败", id)
		return nil
	}
	list := make(map[uint32]*OutfitPreset)
	addOutfit := func(outfitId uint32) bool {
		if outfitId == 0 {
			return true
		}
		return s.GetItemModel().AddItemFashionInfo(outfitId)
	}
	if !addOutfit(uint32(conf.HatID)) {
		goto err
	}
	if !addOutfit(uint32(conf.HairID)) {
		goto err
	}
	if !addOutfit(uint32(conf.ClothID)) {
		goto err
	}
	for i := 0; i < gdconf.GetConstant().OutfitPresetNum; i++ {
		info := &OutfitPreset{
			PresetIndex:            uint32(i),
			Hat:                    uint32(conf.HatID),
			Hair:                   uint32(conf.HairID),
			Clothes:                uint32(conf.ClothID),
			Ornament:               0,
			HatDyeSchemeIndex:      0,
			HairDyeSchemeIndex:     0,
			ClothesDyeSchemeIndex:  0,
			OrnamentDyeSchemeIndex: 0,
			OutfitHideInfo: &OutfitHideInfo{
				HideOrn:   false,
				HideBraid: false,
			},
		}

		list[uint32(i)] = info
	}

	return list
err:
	return nil
}

func (s *Player) GetOutfitPresetList(info *CharacterInfo) map[uint32]*OutfitPreset {
	if info.OutfitPresetList == nil {
		info.OutfitPresetList = s.NewOutfitPresetList(info.CharacterId)
	}
	return info.OutfitPresetList
}

func (s *Player) GetOutfitPreset(info *CharacterInfo, index uint32) *OutfitPreset {
	list := s.GetOutfitPresetList(info)
	return list[index]
}

func (s *Player) GetPbOutfitPresets(characterInfo *CharacterInfo) []*proto.OutfitPreset {
	infoList := s.GetOutfitPresetList(characterInfo)
	if infoList == nil {
		return nil
	}
	pbInfoList := make([]*proto.OutfitPreset, 0)
	for _, info := range infoList {
		pbInfo := &proto.OutfitPreset{
			PresetIndex:                 info.PresetIndex,
			Hat:                         info.Hat,
			Hair:                        info.Hair,
			Clothes:                     info.Clothes,
			Ornament:                    info.Ornament,
			HatDyeSchemeIndex:           info.HatDyeSchemeIndex,
			HairDyeSchemeIndex:          info.HairDyeSchemeIndex,
			ClothesDyeSchemeIndex:       info.ClothesDyeSchemeIndex,
			OrnamentDyeSchemeIndex:      info.OrnamentDyeSchemeIndex,
			OutfitHideInfo:              info.OutfitHideInfo.OutfitHideInfo(),
			PendTop:                     0,
			PendChest:                   0,
			PendPelvis:                  0,
			PendUpFace:                  0,
			PendDownFace:                0,
			PendLeftHand:                0,
			PendRightHand:               0,
			PendLeftFoot:                0,
			PendRightFoot:               0,
			PendTopDyeSchemeIndex:       0,
			PendChestDyeSchemeIndex:     0,
			PendPelvisDyeSchemeIndex:    0,
			PendUpFaceDyeSchemeIndex:    0,
			PendDownFaceDyeSchemeIndex:  0,
			PendLeftHandDyeSchemeIndex:  0,
			PendRightHandDyeSchemeIndex: 0,
			PendLeftFootDyeSchemeIndex:  0,
			PendRightFootDyeSchemeIndex: 0,
		}
		alg.AddList(&pbInfoList, pbInfo)
	}

	return pbInfoList
}

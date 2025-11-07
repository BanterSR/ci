package model

import (
	"gucooing/lolo/gdconf"
	"gucooing/lolo/pkg/log"
	"gucooing/lolo/protocol/proto"
)

type ItemModel struct {
	InstanceIndex   uint32                       // 物品索引生成器
	ItemBadgeMap    map[uint32]*ItemBadgeInfo    // 徽章
	ItemUmbrellaMap map[uint32]*ItemUmbrellaInfo // 伞
	ItemFashionMap  map[uint32]*ItemFashionInfo  // 服装
	ItemWeaponMap   map[uint32]*ItemWeaponInfo   // 武器
}

func DefaultItemModel() *ItemModel {
	return &ItemModel{}
}

func (s *Player) GetItemModel() *ItemModel {
	if s == nil {
		return nil
	}
	if s.Item == nil {
		s.Item = DefaultItemModel()
	}
	return s.Item
}

func (i *ItemModel) NextInstanceIndex() uint32 {
	if i == nil {
		return 0
	}
	i.InstanceIndex++
	return i.InstanceIndex
}

type EBagItemTag interface {
	GetPbItemDetail() *proto.ItemDetail
}

type ItemBadgeInfo struct {
	ItemId uint32
	Num    int64
}

func (i *ItemModel) GetItemBadgeMap() map[uint32]*ItemBadgeInfo {
	if i == nil {
		return nil
	}
	if i.ItemBadgeMap == nil {
		i.ItemBadgeMap = make(map[uint32]*ItemBadgeInfo)
	}
	return i.ItemBadgeMap
}

func (i *ItemModel) AddItemBadge(itemId uint32, num int64) {
	conf := gdconf.GetItemConfigure(itemId)
	list := i.GetItemBadgeMap()
	if conf == nil || list == nil ||
		conf.NewBagItemTag != int32(proto.EBagItemTag_EBagItemTag_Badge) {
		log.Game.Warnf("添加Badge失败,数据异常或不存在ID:%v", itemId)
		return
	}
	info := list[itemId]
	if info == nil {
		info = &ItemBadgeInfo{
			ItemId: itemId,
			Num:    0,
		}
		list[itemId] = info
	}
	info.Num += num
}

func (x *ItemBadgeInfo) GetPbItemDetail() *proto.ItemDetail {
	info := &proto.ItemDetail{
		MainItem: &proto.ItemInfo{
			ItemId:  x.ItemId,
			ItemTag: proto.EBagItemTag_EBagItemTag_Badge,
			Item: &proto.ItemInfo_BaseItem{
				BaseItem: &proto.BaseItem{
					ItemId: x.ItemId,
					Num:    x.Num,
				},
			},
		},
		PackType: proto.ItemDetail_PackType_Inventory,
	}
	return info
}

type ItemUmbrellaInfo struct {
	ItemId uint32
	Num    int64
}

func (i *ItemModel) GetItemUmbrellaMap() map[uint32]*ItemUmbrellaInfo {
	if i == nil {
		return nil
	}
	if i.ItemUmbrellaMap == nil {
		i.ItemUmbrellaMap = make(map[uint32]*ItemUmbrellaInfo)
	}
	return i.ItemUmbrellaMap
}

func (i *ItemModel) AddItemUmbrella(itemId uint32, num int64) {
	conf := gdconf.GetItemConfigure(itemId)
	list := i.GetItemUmbrellaMap()
	if conf == nil || list == nil ||
		conf.NewBagItemTag != int32(proto.EBagItemTag_EBagItemTag_Umbrella) {
		log.Game.Warnf("添加Umbrella失败,数据异常或不存在ID:%v", itemId)
		return
	}
	info := list[itemId]
	if info == nil {
		info = &ItemUmbrellaInfo{
			ItemId: itemId,
			Num:    0,
		}
		list[itemId] = info
	}
	info.Num += num
}

func (i *ItemUmbrellaInfo) GetPbItemDetail() *proto.ItemDetail {
	info := &proto.ItemDetail{
		MainItem: &proto.ItemInfo{
			ItemId:  i.ItemId,
			ItemTag: proto.EBagItemTag_EBagItemTag_Umbrella,
			Item: &proto.ItemInfo_BaseItem{
				BaseItem: &proto.BaseItem{
					ItemId: i.ItemId,
					Num:    i.Num,
				},
			},
		},
		PackType: proto.ItemDetail_PackType_Inventory,
	}
	return info
}

type ItemWeaponInfo struct {
	WeaponId       uint32            // 武器id
	InstanceId     uint32            // 索引id
	Attack         uint32            // 攻击力
	DamageBalance  uint32            // 伤害平衡
	CriticalRatio  uint32            // 临界比率
	RandomProperty []*RandomProperty // 随机属性
	WearerId       uint32            // 装备者id
	Level          uint32            // 等级
	StrengthLevel  uint32            // 强度等级
	StrengthExp    uint32            // 强度经验
	Star           uint32            // 星数
	Inscription1   uint32            // 铭文1
	Durability     uint32            // 耐用性
	PropertyIndex  uint32            // ?指数
	IsLock         bool              // 是否锁
}

type RandomProperty struct {
	PropertyType proto.EPropertyType // 类型
	Value        uint32              // 值
}

func (i *ItemModel) GetItemWeaponMap() map[uint32]*ItemWeaponInfo {
	if i == nil {
		return nil
	}
	if i.ItemWeaponMap == nil {
		i.ItemWeaponMap = make(map[uint32]*ItemWeaponInfo)
	}
	return i.ItemWeaponMap
}

func (i *ItemModel) GetItemWeaponInfo(instanceId uint32) *ItemWeaponInfo {
	list := i.GetItemWeaponMap()
	return list[instanceId]
}

func (i *ItemModel) AddItemWeaponInfo(weaponId uint32) *ItemWeaponInfo {
	conf := gdconf.GetWeaponConfigure(weaponId)
	list := i.GetItemWeaponMap()
	if conf == nil || list == nil {
		log.Game.Warnf("添加Weapon失败,数据异常或不存在ID:%v", weaponId)
		return nil
	}
	instanceId := i.NextInstanceIndex()
	info := &ItemWeaponInfo{
		WeaponId:       uint32(conf.ItemID),
		InstanceId:     instanceId,
		Attack:         1,
		DamageBalance:  1,
		CriticalRatio:  1,
		RandomProperty: make([]*RandomProperty, 0),
		WearerId:       0,
		Level:          1,
		StrengthLevel:  0,
		StrengthExp:    0,
		Star:           1,
		Inscription1:   0,
		Durability:     0,
		PropertyIndex:  1,
		IsLock:         false,
	}
	list[instanceId] = info

	return info
}

func (i *ItemWeaponInfo) GetPbItemDetail() *proto.ItemDetail {
	info := &proto.ItemDetail{
		MainItem: &proto.ItemInfo{
			ItemId:  i.WeaponId,
			ItemTag: proto.EBagItemTag_EBagItemTag_Weapon,
			Item: &proto.ItemInfo_Weapon{
				Weapon: i.GetPbWeaponInstance(),
			},
		},
		PackType: proto.ItemDetail_PackType_Inventory,
	}
	return info
}

func (i *ItemWeaponInfo) SetWearerId(id uint32) {
	i.WearerId = id
}

func (i *ItemWeaponInfo) GetPbWeaponInstance() *proto.WeaponInstance {
	info := &proto.WeaponInstance{
		WeaponId:       i.WeaponId,
		InstanceId:     i.InstanceId,
		Attack:         i.Attack,
		DamageBalance:  i.DamageBalance,
		CriticalRatio:  i.CriticalRatio,
		RandomProperty: make([]*proto.RandomProperty, 0),
		WearerId:       i.WearerId,
		Level:          i.Level,
		StrengthLevel:  i.StrengthLevel,
		StrengthExp:    i.StrengthExp,
		Star:           i.Star,
		Inscription1:   i.Inscription1,
		Durability:     i.Durability,
		PropertyIndex:  i.PropertyIndex,
		IsLock:         i.IsLock,
	}

	return info
}

type ItemFashionInfo struct {
	ItemId uint32
}

func (i *ItemModel) GetItemFashionMap() map[uint32]*ItemFashionInfo {
	if i == nil {
		return nil
	}
	if i.ItemFashionMap == nil {
		i.ItemFashionMap = make(map[uint32]*ItemFashionInfo)
	}
	return i.ItemFashionMap
}

func (i *ItemModel) AddItemFashionInfo(itemId uint32) bool {
	conf := gdconf.GetItemConfigure(itemId)
	list := i.GetItemFashionMap()
	if conf == nil || list == nil ||
		conf.NewBagItemTag != int32(proto.EBagItemTag_EBagItemTag_Fashion) {
		log.Game.Warnf("添加Fashion失败,数据异常或不存在ID:%v", itemId)
		return false
	}
	if list[itemId] != nil {
		return true
	}
	list[itemId] = &ItemFashionInfo{
		ItemId: itemId,
	}
	return true
}

func (i *ItemFashionInfo) GetPbItemDetail() *proto.ItemDetail {
	info := &proto.ItemDetail{
		MainItem: &proto.ItemInfo{
			ItemId:  i.ItemId,
			ItemTag: proto.EBagItemTag_EBagItemTag_Fashion,
			Item: &proto.ItemInfo_Outfit{
				Outfit: &proto.Outfit{
					OutfitId: i.ItemId,
					DyeSchemes: []*proto.OutfitDyeScheme{
						{
							SchemeIndex: 0,
							Colors:      make([]*proto.PosColor, 0),
							IsUnLock:    true,
						},
					},
				},
			},
		},
		PackType: proto.ItemDetail_PackType_Inventory,
	}
	return info
}

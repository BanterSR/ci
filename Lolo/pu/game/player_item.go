package game

import (
	"gucooing/lolo/game/model"
	"gucooing/lolo/pkg/alg"
	"gucooing/lolo/protocol/cmd"
	"gucooing/lolo/protocol/proto"
)

func (g *Game) PackNotice(s *model.Player) {
	notice := &proto.PackNotice{
		Status:          proto.StatusCode_StatusCode_OK,
		Items:           make([]*proto.ItemDetail, 0),
		TempPackMaxSize: 30,
		IsClearTempPack: false,
	}
	defer g.send(s, cmd.PackNotice, 0, notice)
	// 徽章
	for _, v := range s.GetItemModel().GetItemBadgeMap() {
		notice.Items = append(notice.Items, v.GetPbItemDetail())
	}
	// 伞
	for _, v := range s.GetItemModel().GetItemUmbrellaMap() {
		notice.Items = append(notice.Items, v.GetPbItemDetail())
	}
	// 服装
	for _, v := range s.GetItemModel().GetItemFashionMap() {
		notice.Items = append(notice.Items, v.GetPbItemDetail())
	}
	// 武器
	for _, v := range s.GetItemModel().GetItemWeaponMap() {
		notice.Items = append(notice.Items, v.GetPbItemDetail())
	}
}

func (g *Game) GetWeapon(s *model.Player, msg *alg.GameMsg) {
	rsp := &proto.GetWeaponRsp{
		Status:   proto.StatusCode_StatusCode_OK,
		Weapons:  make([]*proto.WeaponInstance, 0),
		TotalNum: uint32(len(s.GetItemModel().GetItemWeaponMap())),
		EndIndex: s.GetItemModel().InstanceIndex,
	}
	defer g.send(s, cmd.GetWeaponRsp, msg.PacketId, rsp)
	for _, v := range s.GetItemModel().GetItemWeaponMap() {
		rsp.Weapons = append(rsp.Weapons, v.GetPbWeaponInstance())
	}
}

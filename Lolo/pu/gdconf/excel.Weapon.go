package gdconf

import (
	"gucooing/lolo/protocol/excel"
)

type Weapon struct {
	AllWeaponDatas *excel.AllWeaponDatas
	WeaponMap      map[int32]*excel.WeaponConfigure
}

func (g *GameConfig) loadWeapon() {
	info := &Weapon{
		AllWeaponDatas: new(excel.AllWeaponDatas),
		WeaponMap:      make(map[int32]*excel.WeaponConfigure),
	}
	g.Excel.Weapon = info
	name := "Weapon.json"
	ReadJson(g.excelPath, name, &info.AllWeaponDatas)
	for _, v := range info.AllWeaponDatas.GetWeapon().GetDatas() {
		info.WeaponMap[v.ID] = v
	}
}

func GetWeaponConfigure(id uint32) *excel.WeaponConfigure {
	return cc.Excel.Weapon.WeaponMap[int32(id)]
}

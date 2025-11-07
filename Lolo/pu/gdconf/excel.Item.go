package gdconf

import (
	"gucooing/lolo/protocol/excel"
)

type Item struct {
	AllItemDatas *excel.AllItemDatas
	ItemMap      map[int32]*excel.ItemConfigure
}

func (g *GameConfig) loadItem() {
	info := &Item{
		AllItemDatas: new(excel.AllItemDatas),
		ItemMap:      make(map[int32]*excel.ItemConfigure),
	}
	g.Excel.Item = info
	name := "Item.json"
	ReadJson(g.excelPath, name, &info.AllItemDatas)
	for _, v := range info.AllItemDatas.GetItem().GetDatas() {
		info.ItemMap[v.ID] = v
	}
}

func GetItemConfigure(id uint32) *excel.ItemConfigure {
	return cc.Excel.Item.ItemMap[int32(id)]
}

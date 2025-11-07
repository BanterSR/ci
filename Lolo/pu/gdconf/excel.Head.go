package gdconf

import (
	"gucooing/lolo/protocol/excel"
)

func (g *GameConfig) loadHead() {
	g.Excel.AllHeadDatas = new(excel.AllHeadDatas)
	name := "Head.json"
	ReadJson(g.excelPath, name, &g.Excel.AllHeadDatas)
}

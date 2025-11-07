package gdconf

import (
	"gucooing/lolo/protocol/excel"
)

type Character struct {
	AllCharacterDatas *excel.AllCharacterDatas
	CharacterMap      map[int32]*excel.CharacterConfigure
}

func (g *GameConfig) loadCharacter() {
	info := &Character{
		AllCharacterDatas: new(excel.AllCharacterDatas),
		CharacterMap:      make(map[int32]*excel.CharacterConfigure),
	}
	g.Excel.Character = info
	name := "Character.json"
	ReadJson(g.excelPath, name, &info.AllCharacterDatas)
	for _, v := range info.AllCharacterDatas.GetCharacter().GetDatas() {
		info.CharacterMap[v.ID] = v
	}
}

func GetCharacterMap() map[int32]*excel.CharacterConfigure {
	return cc.Excel.Character.CharacterMap
}

func GetCharacter(id int32) *excel.CharacterConfigure {
	return cc.Excel.Character.CharacterMap[id]
}

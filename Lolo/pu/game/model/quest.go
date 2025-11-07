package model

import (
	"gucooing/lolo/protocol/proto"
)

func (s *Player) GetQuestDetail() *proto.QuestDetail {
	return &proto.QuestDetail{
		Chapters:                make([]*proto.Chapter, 0),
		DailyQuestBonusDayLeft:  nil,
		DailyQuestBonusWeekLeft: nil,
		RandomQuestBonusLeft:    nil,
		Quests:                  make([]*proto.Quest, 0),
	}
}

func (s *Player) GetPlayerQuestionnaireInfo() *proto.PlayerQuestionnaireInfo {
	return &proto.PlayerQuestionnaireInfo{
		ToFill: make([]*proto.QuestionnaireBrief, 0),
	}
}

package model

import (
	"gucooing/lolo/pkg/log"
	"gucooing/lolo/protocol/proto"
)

type TeamModel struct {
	TeamInfo *TeamInfo
}

type TeamInfo struct {
	Char1 uint32
	Char2 uint32
	Char3 uint32
}

func (s *Player) GetTeamModel() *TeamModel {
	if s.Team == nil {
		s.Team = new(TeamModel)
	}
	return s.Team
}

func (s *Player) newTeamInfo() *TeamInfo {
	info := &TeamInfo{
		Char1: 0,
		Char2: 0,
		Char3: 0,
	}
	index := 0
	for id, _ := range s.GetCharacterMap() {
		switch index {
		case 0:
			info.Char1 = id
		case 1:
			info.Char2 = id
		case 2:
			info.Char3 = id
		}
		index++
		if index == 3 {
			break
		}
	}
	if index != 3 {
		log.Game.Warnf("玩家:%v角色数量不足", s.UserId)
	}
	return info
}

func (s *Player) GetTeamInfo() *TeamInfo {
	info := s.GetTeamModel()
	if info.TeamInfo == nil {
		info.TeamInfo = s.newTeamInfo()
	}
	return info.TeamInfo
}

func (s *Player) GetPbTeam() *proto.Team {
	teamInfo := s.GetTeamInfo()
	info := &proto.Team{
		Char1: teamInfo.Char1,
		Char2: teamInfo.Char2,
		Char3: teamInfo.Char3,
	}
	return info
}

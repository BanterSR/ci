package game

import (
	"encoding/json"
	"strconv"
	"time"

	"gucooing/lolo/db"
	"gucooing/lolo/game/model"
	"gucooing/lolo/pkg/alg"
	"gucooing/lolo/pkg/log"
	"gucooing/lolo/pkg/ofnet"
	"gucooing/lolo/protocol/cmd"
	"gucooing/lolo/protocol/proto"
)

func (g *Game) PlayerLogin(conn ofnet.Conn, userId uint32, msg *alg.GameMsg) {
	req := msg.Body.(*proto.PlayerLoginReq)
	rsp := &proto.PlayerLoginRsp{
		Status:           proto.StatusCode_StatusCode_OK,
		IsReconnect:      req.IsReconnect, // 是否重新连接
		ReconnectSuccess: req.IsReconnect, // 重新连接是否成功
	}
	// 重复登录检查
	if player := g.GetUser(userId); player != nil {
		g.kickPlayer(userId) // 下线老玩家
	}
	// 拉取数据
	dbUser, err := db.GetOFGameByUserId(userId)
	if err != nil {
		rsp.Status = proto.StatusCode_StatusCode_PLAYER_NOT_FOUND
		log.Game.Warnf("数据库拉取玩家:%v数据失败:%s", userId, err.Error())
		return
	}
	s := &model.Player{
		UserId:    userId,
		Conn:      conn,
		Online:    true,
		NetFreeze: false,
		Created:   dbUser.CreatedAt,
		Updated:   dbUser.UpdatedAt,
	}
	if dbUser.BinData != nil {
		if err := json.Unmarshal(dbUser.BinData, s); err != nil {
			rsp.Status = proto.StatusCode_StatusCode_PLAYER_NOT_FOUND
			log.Game.Warnf("玩家:%v数据序列化失败:%s", userId, err.Error())
			return
		}
	} else {
		// newPlayer
	}
	g.userMap[userId] = s

	basic := s.GetBasicModel()
	if basic == nil {
		log.Game.Warnf("UserId:%v 登录失败,玩家数据异常", s.UserId)
		rsp.Status = proto.StatusCode_StatusCode_PLAYER_NOT_FOUND
		return
	}
	defer func() {
		g.send(s, cmd.PlayerLoginRsp, msg.PacketId, rsp)
		if req.IsReconnect {
			g.loginGame(s)
		}
	}()
	// pack
	{
		rsp.ClientSeqId = msg.SeqId
		rsp.ServerSeqId = s.GetSeqId()
	}
	// 基础信息
	{
		rsp.PlayerName = basic.PlayerName
		rsp.RegisterTime = uint32(s.Created.Unix())
		rsp.AnalysisAccountId = strconv.Itoa(int(s.UserId))
	}
	// 加入房间
	scenePlayer := g.getWordInfo().addScenePlayer(s)
	if scenePlayer == nil {
		rsp.Status = proto.StatusCode_StatusCode_SCENE_CHANNEL_IS_FULL
		return
	}
	// 场景
	{
		rsp.SceneId = scenePlayer.SceneId
		rsp.ChannelId = scenePlayer.ChannelId
	}
	// 其他信息
	{
		rsp.ServerTimeMs = time.Now().UnixMilli()
		rsp.RegionName = "cn_prod_main"
		rsp.PlayerAgeRange = 0
		rsp.Tags = 0
		rsp.ServerTimeZone = 28800
		rsp.ClientLogServerToken = "114514"
	}
	log.Game.Infof("UserId:%v Name:%s IP:%s 设备:%s 系统:%s 登录成功!",
		s.UserId, rsp.PlayerName, req.Ip, req.DeviceModel, req.OsVer)
}

func (g *Game) PlayerMainData(s *model.Player, msg *alg.GameMsg) {
	rsp := &proto.PlayerMainDataRsp{
		Status: proto.StatusCode_StatusCode_OK,
	}
	defer func() {
		g.send(s, cmd.PlayerMainDataRsp, msg.PacketId, rsp)
		g.loginGame(s)
	}()
	// 基础信息
	basic := s.GetBasicModel()
	{
		rsp.PlayerId = s.UserId
		rsp.PlayerLabel = s.UserId // 玩家标签
		rsp.PlayerName = basic.PlayerName
		rsp.Level = basic.Level
		rsp.Sign = basic.Sign
		rsp.Exp = basic.Exp
		rsp.Head = basic.Head
		rsp.CreateTime = uint32(s.Created.Unix())
		rsp.Birthday = basic.Birthday
		rsp.IsHideBirthday = basic.IsHideBirthday
		rsp.PhoneBackground = basic.PhoneBackground
		rsp.Appearance = &proto.PlayerAppearance{
			AvatarFrame: 0,
			Pendant:     0,
		}
	}
	// 已获得的角色
	{
		rsp.Characters = s.GetAllPbCharacter()
	}
	scenePlayer := g.getWordInfo().getScenePlayer(s)
	if scenePlayer == nil {
		rsp.Status = proto.StatusCode_StatusCode_SCENE_CHANNEL_IS_FULL
		return
	}
	// 场景
	{
		rsp.ChannelId = scenePlayer.ChannelId
		rsp.ChannelLabel = scenePlayer.ChannelId // 房间标签
		rsp.SceneId = scenePlayer.SceneId
	}
	// 队伍
	{
		rsp.Team = s.GetPbTeam()
	}
	// buff
	{
		rsp.PlayerBuffs = make([]*proto.PlayerBuff, 0)
	}
	// 其他
	{
		rsp.PlayerDropRateInfo = s.GetPbPlayerDropRateInfo()
		rsp.QuestDetail = s.GetQuestDetail()
		rsp.QuestionnaireInfo = s.GetPlayerQuestionnaireInfo()
		rsp.UnlockFunctions = make([]uint32, 0)
	}
}

func (g *Game) loginGame(s *model.Player) {
	g.PackNotice(s)
	// 进入房间
	g.joinSceneChannel(s)
	/*
		s.ChangeChatChannel()

	*/
	g.ChatMsgRecordInitNotice(s)
}

func (g *Game) PlayerPing(s *model.Player, msg *alg.GameMsg) {
	req := msg.Body.(*proto.PlayerPingReq)
	rsp := &proto.PlayerPingRsp{
		Status:       proto.StatusCode_StatusCode_OK,
		ClientTimeMs: req.ClientTimeMs,
		ServerTimeMs: time.Now().UnixMilli(),
	}
	defer g.send(s, cmd.PlayerPingRsp, msg.PacketId, rsp)
}

func (g *Game) GamePlayReward(s *model.Player, msg *alg.GameMsg) {
	// req := msg.Body.(*proto.GamePlayRewardReq)
	rsp := &proto.GamePlayRewardRsp{
		Status:                 proto.StatusCode_StatusCode_OK,
		DynamicTreasureBoxInfo: nil,
		Items:                  make([]*proto.ItemDetail, 0),
	}
	defer g.send(s, cmd.GamePlayRewardRsp, msg.PacketId, rsp)
}

func (g *Game) AcceptQuest(s *model.Player, msg *alg.GameMsg) {
	req := msg.Body.(*proto.AcceptQuestReq)
	rsp := &proto.AcceptQuestRsp{
		Status: proto.StatusCode_StatusCode_OK,
		Quest: &proto.Quest{
			QuestId:       req.QuestId,
			Conditions:    make([]*proto.Condition, 0),
			Status:        proto.QuestStatus_QuestStatus_InProgress,
			CompleteCount: 0,
			BonusTimes:    0,
			ActivityId:    0,
		},
	}
	defer g.send(s, cmd.AcceptQuestRsp, msg.PacketId, rsp)
}

func (g *Game) GetAchieveOneGroup(s *model.Player, msg *alg.GameMsg) {
	// req := msg.Body.(*proto.GetAchieveOneGroupReq)
	rsp := &proto.GetAchieveOneGroupRsp{
		Status: proto.StatusCode_StatusCode_OK,
		CurrentGroupAchieveInfo: &proto.OneGroupAchieveInfo{
			GroupId:              0,
			RewardedAchieveIdLst: make([]uint32, 0),
			AchieveLst:           make([]*proto.Achieve, 0),
			FinishAchieveLst:     make([]*proto.FinishAchieveInfo, 0),
		},
		IsReward: false,
	}
	defer g.send(s, cmd.GetAchieveOneGroupRsp, msg.PacketId, rsp)
}

func (g *Game) GetAchieveGroupList(s *model.Player, msg *alg.GameMsg) {
	// req := msg.Body.(*proto.GetAchieveGroupListReq)
	rsp := &proto.GetAchieveGroupListRsp{
		Status:             proto.StatusCode_StatusCode_OK,
		RewardedGroupIdLst: make([]uint32, 0),
	}
	defer g.send(s, cmd.GetAchieveGroupListRsp, msg.PacketId, rsp)
}

func (g *Game) GenericGameB(s *model.Player, msg *alg.GameMsg) {
	// req := msg.Body.(*proto.GenericGameBReq)
	rsp := &proto.GenericGameBRsp{
		Status:       proto.StatusCode_StatusCode_OK,
		GenericMsgId: 0,
		Params:       make([]*proto.CommonParam, 0),
	}
	defer g.send(s, cmd.GenericGameBRsp, msg.PacketId, rsp)
}

func (g *Game) GetCollectItemIds(s *model.Player, msg *alg.GameMsg) {
	// req := msg.Body.(*proto.GetCollectItemIdsReq)
	rsp := &proto.GetCollectItemIdsRsp{
		Status:  proto.StatusCode_StatusCode_OK,
		ItemIds: make([]uint32, 0),
	}
	defer g.send(s, cmd.GetCollectItemIdsRsp, msg.PacketId, rsp)
}

func (g *Game) ManualList(s *model.Player, msg *alg.GameMsg) {
	// req := msg.Body.(*proto.ManualListReq)
	rsp := &proto.ManualListRsp{
		Status: proto.StatusCode_StatusCode_OK,
		Flags:  make([]*proto.ManualFlag, 0),
	}
	defer g.send(s, cmd.ManualListRsp, msg.PacketId, rsp)
}

func (g *Game) GetCollectMoonInfo(s *model.Player, msg *alg.GameMsg) {
	req := msg.Body.(*proto.GetCollectMoonInfoReq)
	rsp := &proto.GetCollectMoonInfoRsp{
		Status:           proto.StatusCode_StatusCode_OK,
		SceneId:          req.SceneId,
		CollectedMoonIds: make([]uint32, 0),
		EmotionMoons:     make([]*proto.EmotionMoonInfo, 0),
	}
	defer g.send(s, cmd.GetCollectMoonInfoRsp, msg.PacketId, rsp)
}

func (g *Game) ChangeMusicalItem(s *model.Player, msg *alg.GameMsg) {
	// req := msg.Body.(*proto.ChangeMusicalItemReq)
	rsp := &proto.ChangeMusicalItemRsp{
		Status:                proto.StatusCode_StatusCode_OK,
		Source:                0,
		MusicalItemInstanceId: 0,
		MusicalItemId:         0,
	}
	defer g.send(s, cmd.ChangeMusicalItemRsp, msg.PacketId, rsp)
}

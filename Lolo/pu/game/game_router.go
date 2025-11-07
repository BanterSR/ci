package game

import (
	"gucooing/lolo/game/model"
	"gucooing/lolo/pkg/alg"
	"gucooing/lolo/pkg/log"
	"gucooing/lolo/pkg/ofnet"
	"gucooing/lolo/protocol/cmd"
)

type HandlerFunc func(s *model.Player, msg *alg.GameMsg)

func (g *Game) newRouter() {
	g.handlerFuncRouteMap = map[uint32]HandlerFunc{
		// cmd.PlayerLoginReq: g.PlayerLogin 登录第二个包 仅提示用
		// 基础包
		cmd.PlayerPingReq: g.PlayerPing, // ping
		// 玩家基础信息
		cmd.PlayerMainDataReq: g.PlayerMainData, // 获取玩家信息
		// 场景
		cmd.PlayerSceneRecordReq: g.PlayerSceneRecord, // 玩家场景同步器
		cmd.SendActionReq:        g.SendAction,        // 场景自动化同步器
		// 队伍
		cmd.UpdateTeamReq: g.UpdateTeam, // 更新队伍
		// 武器
		cmd.GetWeaponReq: g.GetWeapon, // 获取玩家全部武器
		// 卡池
		cmd.GachaListReq: g.GachaList, // 获取卡池信息

		cmd.GenericSceneBReq:          g.GenericSceneB,
		cmd.AbilityBadgeListReq:       g.AbilityBadgeList,
		cmd.SceneProcessListReq:       g.SceneProcessList,
		cmd.ShopInfoReq:               g.ShopInfo,
		cmd.FriendReq:                 g.Friend,
		cmd.WishListByFriendIdReq:     g.WishListByFriendId,
		cmd.GetLifeInfoReq:            g.GetLifeInfo,
		cmd.GetMailsReq:               g.GetMails,
		cmd.GetAchieveOneGroupReq:     g.GetAchieveOneGroup,
		cmd.GetAchieveGroupListReq:    g.GetAchieveGroupList,
		cmd.GenericGameBReq:           g.GenericGameB,
		cmd.GetCollectItemIdsReq:      g.GetCollectItemIds,
		cmd.ManualListReq:             g.ManualList,
		cmd.GetCollectMoonInfoReq:     g.GetCollectMoonInfo,
		cmd.ChangeMusicalItemReq:      g.ChangeMusicalItem,
		cmd.GetArchiveInfoReq:         g.GetArchiveInfo,         //
		cmd.PlayerAbilityListReq:      g.PlayerAbilityList,      //
		cmd.PosterIllustrationListReq: g.PosterIllustrationList, //
		cmd.WorldLevelAchieveListReq:  g.WorldLevelAchieveList,  //
		cmd.SupplyBoxInfoReq:          g.SupplyBoxInfo,          //
		cmd.GetAllCharacterEquipReq:   g.GetAllCharacterEquip,   //
		cmd.GamePlayRewardReq:         g.GamePlayReward,         //
		cmd.AcceptQuestReq:            g.AcceptQuest,            //
	}
}

func (g *Game) RouteHandle(conn ofnet.Conn, userId uint32, msg *alg.GameMsg) {
	if msg.MsgId == cmd.PlayerLoginReq {
		g.PlayerLogin(conn, userId, msg)
		return
	}
	handlerFunc, ok := g.handlerFuncRouteMap[msg.MsgId]
	if !ok {
		log.Game.Errorf("no route for msg, cmdId: %v name:%s", msg.MsgId, cmd.Get().GetCmdNameByCmdId(msg.MsgId))
		return
	}
	player := g.GetUser(userId)
	if player == nil {
		log.Game.Errorf("player is nil, userId: %v", userId)
		return
	}
	if !player.Online {
		log.Game.Errorf("player not online, userId: %v", userId)
		return
	}
	if player.NetFreeze {
		return
	}
	handlerFunc(player, msg)
}

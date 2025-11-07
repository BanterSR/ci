package game

import (
	"runtime"

	pb "google.golang.org/protobuf/proto"

	"gucooing/lolo/config"
	"gucooing/lolo/game/model"
	"gucooing/lolo/pkg/alg"
	"gucooing/lolo/pkg/log"
	"gucooing/lolo/pkg/ofnet"
)

type Game struct {
	gameMsgChan         chan *GameMsg
	userMap             map[uint32]*model.Player
	handlerFuncRouteMap map[uint32]HandlerFunc
	wordInfo            *WordInfo
}

type GameMsg struct {
	UserId uint32
	Conn   ofnet.Conn
	*alg.GameMsg
}

func NewGame() *Game {
	conf := config.GetGame()
	log.NewGame()
	g := &Game{
		gameMsgChan: make(chan *GameMsg, conf.MsgChanSize),
		userMap:     make(map[uint32]*model.Player),
	}
	g.newRouter()

	go g.gameMainLoop()
	return g
}

// 游戏主线程
func (g *Game) gameMainLoop() {
	runtime.LockOSThread()
	defer func() {
		runtime.UnlockOSThread()
		if err := recover(); err != nil {
			log.Game.Error("!!! GAME MAIN LOOP PANIC !!!")
			log.Game.Error("error: %v", err)
		}
	}()
	for {
		select {
		case msg := <-g.gameMsgChan:
			g.RouteHandle(msg.Conn, msg.UserId, msg.GameMsg)
		}
	}
}

func (g *Game) send(s *model.Player, cmdId uint32, packetId uint32, payloadMsg pb.Message) {
	if s.NetFreeze {
		return
	}
	s.Conn.Send(cmdId, packetId, payloadMsg)
}

func (g *Game) GetUser(userId uint32) *model.Player {
	player, ok := g.userMap[userId]
	if !ok {
		return nil
	}
	return player
}

func (g *Game) kickPlayer(userId uint32) {}

func (g *Game) GetGameMsgChan() chan *GameMsg {
	return g.gameMsgChan
}

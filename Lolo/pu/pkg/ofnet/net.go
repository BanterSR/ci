package ofnet

import (
	"errors"
	"net"

	"google.golang.org/protobuf/encoding/protojson"
	pb "google.golang.org/protobuf/proto"

	"gucooing/lolo/pkg/alg"
	"gucooing/lolo/pkg/log"
	"gucooing/lolo/protocol/cmd"
)

var CLIENT_CONN_NUM int64 = 0 // 当前客户端连接数
var TPS int64

type Net interface {
	Accept() (Conn, error)
	Close() error
	SetBlackPackId(packIdList map[uint32]struct{})
	SetLogMsg(logMsg bool)
}

func NewNet(network, addr string) (Net, error) {
	log.Gate.Infof("协议:%s,启动在%s 上", network, addr)
	switch network {
	case "tcp":
		return newTcpNet(addr)
	}
	return nil, errors.New("network not support")
}

type netBase struct {
	logMsg      bool
	blackPackId map[uint32]struct{}
}

func (c *netBase) SetBlackPackId(packIdList map[uint32]struct{}) {
	c.blackPackId = packIdList
}

func (c *netBase) SetLogMsg(logMsg bool) {
	c.logMsg = logMsg
}

func (c *netBase) logPack(packId uint32) bool {
	if !c.logMsg {
		return false
	}
	_, ok := c.blackPackId[packId]
	return !ok
}

type Conn interface {
	Read() (*alg.GameMsg, error)
	Send(cmdId, packetId uint32, protoObj pb.Message)
	SetUID(uint32)
	GetSeqId() uint32
	SetServerTag(serverTag string)
	Close()
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
}

const (
	ClientMsg = iota
	ServerMsg
)

func logMag(tp int, serverTag string, isLog bool, uid, cmdId uint32, payloadMsg pb.Message) {
	if !isLog {
		return
	}
	var s string
	switch tp {
	case ClientMsg:
		s = "c -> s"
	case ServerMsg:
		s = "s -> c"
	}
	log.Gate.Debugf("%s[Server:%s][UID:%v][CMD:%s]Pack:%s",
		s,
		serverTag,
		uid,
		cmd.Get().GetCmdNameByCmdId(cmdId),
		protojson.Format(payloadMsg),
	)
}

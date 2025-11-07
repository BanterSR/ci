package ofnet

import (
	"bufio"
	"encoding/binary"
	"errors"
	"io"
	"net"

	"github.com/golang/snappy"
	pb "google.golang.org/protobuf/proto"

	"gucooing/lolo/pkg/alg"
	"gucooing/lolo/pkg/log"
	"gucooing/lolo/protocol/cmd"
	"gucooing/lolo/protocol/proto"
)

type tcpNet struct {
	*netBase
	listener net.Listener
}

func newTcpNet(addr string) (*tcpNet, error) {
	x := &tcpNet{
		netBase: &netBase{
			blackPackId: make(map[uint32]struct{}),
		},
	}
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	x.listener = listener
	return x, nil
}

func (x *tcpNet) Accept() (Conn, error) {
	if x == nil {
		return nil, errors.New("tcpNet is nil")
	}
	conn, err := x.listener.Accept()
	if err != nil {
		return nil, err
	}
	tconn := &tcpConn{
		net:  x,
		conn: conn,
		buf:  bufio.NewReaderSize(conn, alg.PacketMaxLen),
	}

	return tconn, nil
}

func (x *tcpNet) Close() error {
	if x == nil {
		return nil
	}
	return x.listener.Close()
}

type tcpConn struct {
	net       *tcpNet
	conn      net.Conn
	buf       *bufio.Reader
	uid       uint32
	seqId     uint32
	serverTag string
}

func (x *tcpConn) GetSeqId() uint32 {
	return x.seqId
}

func (x *tcpConn) Read() (*alg.GameMsg, error) {
	for {
		// head
		headLenByte := make([]byte, alg.TcpHeadSize)
		_, err := x.buf.Read(headLenByte)
		if err != nil {
			return nil, err
		}
		headLen := binary.BigEndian.Uint16(headLenByte)

		headByte := make([]byte, headLen)
		_, err = x.buf.Read(headByte)
		if err != nil {
			return nil, err
		}
		head := new(proto.PacketHead)
		err = pb.Unmarshal(headByte, head)
		if err != nil {
			log.Gate.Errorf("Could not parse PacketHead proto Error:%s\n", err)
			return nil, err
		}

		// body
		bodyByte := make([]byte, head.BodyLen)
		_, err = x.buf.Read(bodyByte)
		if err != nil {
			return nil, err
		}
		bodyByte = alg.HandleFlag(head.Flag, bodyByte)
		protoObj := cmd.Get().GetProtoObjByCmdId(head.MsgId)
		if protoObj == nil {
			log.Gate.Errorf("protoObj by cmdId:%d\n", head.MsgId)
			continue
		}
		err = pb.Unmarshal(bodyByte, protoObj)
		if err != nil {
			log.Gate.Errorf("unmarshal proto data err: %v\n", err)
			return nil, err
		}
		logMag(ClientMsg,
			x.serverTag,
			x.net.logPack(head.MsgId),
			x.uid,
			head.MsgId,
			protoObj)
		gameMsg := &alg.GameMsg{
			PacketHead: head,
			Body:       protoObj,
		}

		return gameMsg, nil
	}
}

func (x *tcpConn) Send(cmdId, packetId uint32, protoObj pb.Message) {
	if x == nil {
		return
	}

	bodyByte, err := pb.Marshal(protoObj)
	if err != nil {
		log.Gate.Errorf("marshal proto data err: %v\n", err)
		return
	}

	logMag(ServerMsg, x.serverTag, x.net.logPack(cmdId), x.uid, cmdId, protoObj)

	head := &proto.PacketHead{
		MsgId:    cmdId,
		Flag:     0,
		BodyLen:  0,
		SeqId:    x.seqId,
		PacketId: packetId,

		TotalPackCount: 0,
	}
	x.seqId++

	if len(bodyByte) > alg.SnappySize {
		bodyByte = snappy.Encode(nil, bodyByte)
		head.Flag = 1
	}
	head.BodyLen = uint32(len(bodyByte))
	headBytes, err := pb.Marshal(head)
	if err != nil {
		log.Gate.Errorf("marshal proto data err: %v\n", err)
		return
	}
	bin := make([]byte, alg.TcpHeadSize+len(headBytes)+len(bodyByte))

	binary.BigEndian.PutUint16(bin[:alg.TcpHeadSize], uint16(len(headBytes)))
	// 头部数据
	copy(bin[alg.TcpHeadSize:], headBytes)
	// proto数据
	copy(bin[alg.TcpHeadSize+len(headBytes):], bodyByte)

	_, err = x.conn.Write(bin)
	if err != nil && !errors.Is(err, io.ErrClosedPipe) {
		log.Gate.Errorf("tcpConn write error: %v", err)
		return
	}
}

func (x *tcpConn) SetUID(uid uint32) {
	x.uid = uid
}

func (x *tcpConn) SetServerTag(serverTag string) {
	x.serverTag = serverTag
}

func (x *tcpConn) Close() {
	x.conn.Close()
}

func (x *tcpConn) LocalAddr() net.Addr {
	return x.conn.LocalAddr()
}

func (x *tcpConn) RemoteAddr() net.Addr {
	return x.conn.RemoteAddr()
}

package alg

import (
	"log"

	"github.com/golang/snappy"
	pb "google.golang.org/protobuf/proto"

	"gucooing/lolo/protocol/proto"
)

const (
	TcpHeadSize  = 2
	PacketMaxLen = 512000 // 最大应用层包长度
	SnappySize   = 1 << 10
)

type GameMsg struct {
	*proto.PacketHead
	Body pb.Message
}

func HandleFlag(flag uint32, body []byte) []byte {
	switch flag {
	case 0:
		// 不处理
		return body
	case 1:
		var dst []byte
		dst, _ = snappy.Decode(nil, body)
		return dst
	default:
		log.Printf("Unknown flag:%d\n", flag)
		return body
	}
}

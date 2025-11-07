package gateway

import (
	"errors"
	"io"

	"gucooing/lolo/db"
	"gucooing/lolo/game"
	"gucooing/lolo/pkg/alg"
	"gucooing/lolo/pkg/log"
	"gucooing/lolo/pkg/ofnet"
	"gucooing/lolo/protocol/cmd"
	"gucooing/lolo/protocol/proto"
)

type LoginInfo struct {
	*proto.VerifyLoginTokenReq
	conn ofnet.Conn
}

func (g *Gateway) loginSessionManagement() {
	loginMap := make(map[string]*LoginInfo)
	for {
		select {
		case login := <-g.loginChan:
			if _, ok := loginMap[login.SdkUid]; ok {
				// 重复的登录请求
				continue
			}
			loginMap[login.SdkUid] = login
			g.VerifyLoginToken(login)
		case sdkUid := <-g.delLoginChan:
			delete(loginMap, sdkUid)
		}
	}
}

func (g *Gateway) VerifyLoginToken(req *LoginInfo) {
	// 由于没有sdk 所以这里同意全部登录请求
	rsp := &proto.VerifyLoginTokenRsp{
		AccountType: req.AccountType,
		SdkUid:      req.SdkUid,
		DeviceUuid:  req.DeviceUuid,

		Status:       0,
		UserId:       0,
		IsServerOpen: false,
		Text:         "",
		TimeLeft:     0,
		Os:           0,
		BanEndTime:   0,
	}
	defer func() {
		req.conn.Send(cmd.VerifyLoginTokenRsp, 0, rsp)
		g.delLoginChan <- req.SdkUid
	}()
	sdkUid := alg.S2U32(req.SdkUid)
	// token验证
	ofUser, err := db.GetOFUserBySdkUid(sdkUid)
	if err != nil {
		rsp.Status = proto.StatusCode_StatusCode_ACCOUNT_UNAUTH
		log.Gate.Debugf("SdkUid:%s,拉取账号失败err:%s", req.SdkUid, err.Error())
		return
	}

	// TODO 代办工作
	// 验证是否被ban

	// 检查在线满了?

	rsp.IsServerOpen = true
	rsp.Status = proto.StatusCode_StatusCode_OK
	rsp.TimeLeft = 4294967295
	rsp.UserId = ofUser.UserId

	req.conn.SetUID(ofUser.UserId)
	log.Gate.Infof("UserId:%v 平台:%s 正在登录中...", ofUser.UserId, proto.AccountType(req.AccountType).String())

	go g.receive(req.conn, ofUser.UserId)
}

func (g *Gateway) receive(conn ofnet.Conn, userId uint32) {
	for {
		msg, err := conn.Read()
		switch {
		case err == nil:
			g.game.GetGameMsgChan() <- &game.GameMsg{
				UserId:  userId,
				Conn:    conn,
				GameMsg: msg,
			}
		case errors.Is(err, io.EOF),
			errors.Is(err, io.ErrClosedPipe):
			return
		default:
			log.Gate.Errorf("%s", err)
			return
		}
	}
}

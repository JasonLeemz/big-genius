package controllers

import (
	"big-genius/core/errors"
	"big-genius/core/log"
	"big-genius/core/utils"
	"big-genius/core/utils/wechat"
	"big-genius/internal/app/services/mq"
	"github.com/kataras/iris/v12/context"
)

type RobotReq struct {
	MsgSignature string `form:"msg_signature"`
	Timestamp    string `form:"timestamp"`
	Nonce        string `form:"nonce"`
	EchoStr      string `form:"echostr"`
}

func AIRobot(ctx *context.Context) {
	//http://api.3dept.com/?msg_signature=ASDFQWEXZCVAQFASDFASDFSS&timestamp=13500001234&nonce=123412323&echostr=ENCRYPT_STR
	req := RobotReq{}
	err := ctx.ReadForm(&req)

	if err != nil {
		utils.Reply(ctx, nil, errors.GenErr(err))
	}

	if req.EchoStr != "" {
		// signature, timestamp, nonce, EchoStr
		sEchoStr, err := wechat.VerifyUrl(req.MsgSignature, req.Timestamp, req.Nonce, req.EchoStr)
		if err != nil {
			utils.Reply(ctx, sEchoStr, errors.GenErr(err))
		}

		ctx.Write([]byte(sEchoStr))
	}

	body, err := ctx.GetBody()
	log.Logger.Infof("wxsend:%s", body)
	sendMsg(ctx, req, body)
}

func sendMsg(ctx *context.Context, req RobotReq, body []byte) {

	// reqMsgSign, reqTimestamp, reqNonce string, reqData []byte
	userMsg, err := wechat.VerifyData(req.MsgSignature, req.Timestamp, req.Nonce, body)
	if err != nil {
		log.Logger.Errorf("VerifyData error:%s", err.Error())
		return
	}

	msgCont := wechat.MsgContent{
		ToUsername:   userMsg.FromUsername,
		FromUsername: userMsg.ToUsername,
		CreateTime:   userMsg.CreateTime,
		MsgType:      userMsg.MsgType,
		Msgid:        userMsg.Msgid,
		Agentid:      userMsg.Agentid,
	}

	// webhook响应不返回有效数据 msgCont.Content = ""
	msg, err := wechat.GenWebhookMsg(msgCont, req.Timestamp, req.Nonce)
	if err != nil {
		log.Logger.Errorf("GenSendMsg error:%s", err.Error())
		return
	}
	ctx.Write(msg)

	msgCont.Content = userMsg.Content
	mq.SendMQMsg(msgCont)
}

// 如果是重发消息，获取锁后不进行处理
//func getLock(ctx ctx.Context, key string) (int, error) {
//	value := "lock"
//
//	lock, err := logic.RedisGet(ctx, key)
//	if err != nil {
//		log.Logger.Errorf("getLock Error: %s", err.Error())
//		return -1, err
//	}
//
//	if lock == redis.KeyNotExist {
//		// 执行加锁
//
//		l, e := logic.RedisLock(ctx, key, value, 10*time.Second)
//		if e != nil {
//			log.Logger.Errorf("logic.RedisLock Error: %s", e.Error())
//			return -2, e
//		}
//
//		if l {
//			// 如果是本次上锁，返回0
//			return 0, nil
//		}
//	}
//
//	if lock == value {
//		return 1, nil
//	}
//
//	log.Logger.Warnf("lock may be not set,key:%s,value:%s", key, lock)
//	return -3, nil
//}

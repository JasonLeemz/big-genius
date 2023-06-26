package controllers

import (
	ctx "big-genius/core/context"
	"big-genius/core/errors"
	"big-genius/core/log"
	"big-genius/core/utils/wechat"
	"big-genius/internal/app/logic"
	"big-genius/internal/app/services"
	"sync"
	"time"
)

type RobotReq struct {
	MsgSignature string `form:"msg_signature"`
	Timestamp    string `form:"timestamp"`
	Nonce        string `form:"nonce"`
	EchoStr      string `form:"echostr"`
}

func AIRobot(ctx ctx.Context) {
	//http://api.3dept.com/?msg_signature=ASDFQWEXZCVAQFASDFASDFSS&timestamp=13500001234&nonce=123412323&echostr=ENCRYPT_STR
	req := RobotReq{}
	err := ctx.ReadForm(&req)
	log.Logger.Infof("wxform:%v", req)

	if err != nil {
		ctx.Reply(nil, errors.GenErr(err))
	}

	if req.EchoStr != "" {
		// signature, timestamp, nonce, EchoStr
		sEchoStr, err := wechat.VerifyUrl(req.MsgSignature, req.Timestamp, req.Nonce, req.EchoStr)
		if err != nil {
			ctx.Reply(sEchoStr, errors.GenErr(err))
		}

		ctx.Write([]byte(sEchoStr))
	}

	body, err := ctx.GetBody()
	log.Logger.Infof("wxsend:%s", body)
	sendMsg(ctx, req, body)
}

func sendMsg(ctx ctx.Context, req RobotReq, body []byte) {

	// reqMsgSign, reqTimestamp, reqNonce string, reqData []byte
	userMsg, err := wechat.VerifyData(req.MsgSignature, req.Timestamp, req.Nonce, body)
	if err != nil {
		log.Logger.Errorf("VerifyData error:%s", err.Error())
		return
	}
	log.Logger.Infof("userMsg:%v,%v", userMsg, err)

	msgCont := wechat.MsgContent{
		ToUsername:   userMsg.FromUsername,
		FromUsername: userMsg.ToUsername,
		CreateTime:   userMsg.CreateTime,
		MsgType:      userMsg.MsgType,
		Msgid:        userMsg.Msgid,
		Agentid:      userMsg.Agentid,
	}

	log.Logger.Info("=============================请求openai:开始=============================")
	wg := sync.WaitGroup{}
	wg.Add(1)
	// 请求chatgpt
	s := services.NewOpenAIService()
	go func() {
		log.Logger.Info("=============================线程=============================")

		lock, e := getLock(ctx, msgCont.Msgid)
		if e != nil {
			log.Logger.Errorf("getLock Error:%s", e.Error())
		}
		if lock == 0 {
			answer, err := s.CreateChatCompletion(ctx, userMsg.Content)
			if err != nil {
				log.Logger.Errorf("请求openai超时:%s", err.Error())
				msgCont.Content = "请求openai超时"
			} else {
				log.Logger.Errorf("answer:%s", answer)
				msgCont.Content = answer
			}
		}

		log.Logger.Infof("getLock info:%d", lock)
		ctx.Write(nil)
		wg.Done()
	}()

	wg.Wait()
	log.Logger.Info("=============================请求openai:结束=============================")
	if msgCont.Content == "" {
		msgCont.Content = errors.GetErrMsg(errors.ErrNoUnknownInternal)
	}
	msg, err := wechat.GenSendMsg(msgCont, req.Timestamp, req.Nonce)
	if err != nil {
		log.Logger.Errorf("GenSendMsg error:%s", err.Error())
		return
	}

	ctx.Write(msg)
}

// 如果是重发消息，获取锁后不进行处理
func getLock(ctx ctx.Context, key string) (int, error) {
	value := "lock"

	lock, err := logic.RedisGet(ctx, key)
	if err != nil {
		log.Logger.Errorf("getLock Error: %s", err.Error())
		return -1, err
	}

	if lock == logic.KeyNotExist {
		// 执行加锁

		l, e := logic.RedisLock(ctx, key, value, 10*time.Second)
		if e != nil {
			log.Logger.Errorf("logic.RedisLock Error: %s", e.Error())
			return -2, e
		}

		if l {
			// 如果是本次上锁，返回0
			return 0, nil
		}
	}

	if lock == value {
		return 1, nil
	}

	log.Logger.Warnf("lock may be not set,key:%s,value:%s", key, lock)
	return -3, nil
}

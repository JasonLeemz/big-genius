package controllers

import (
	ctx "big-genius/core/context"
	"big-genius/core/errors"
	"big-genius/core/log"
	"big-genius/core/utils/wechat"
	"big-genius/internal/app/services"
)

type RobotReq struct {
	MsgSignature string `form:"msg_signature"`
	Timestamp    string `form:"timestamp"`
	Nonce        string `form:"nonce"`
	EchoStr      string `form:"echostr"`
}

func AIRobot(ctx ctx.Context) {
	//http://api.3dept.com/?msg_signature=ASDFQWEXZCVAQFASDFASDFSS&timestamp=13500001234&nonce=123412323&echostr=ENCRYPT_STR
	//f := ctx.Params()
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
		Msgid:        "",
		Agentid:      userMsg.Agentid,
	}

	// 请求chatgpt
	s := services.NewOpenAIService()
	answer, err := s.CreateChatCompletion(ctx, userMsg.Content)
	if err != nil {
		log.Logger.Errorf("请求openai超时:%s", err.Error())
		msgCont.Content = "请求openai超时"
	} else {
		msgCont.Content = answer
	}

	msg, err := wechat.GenSendMsg(msgCont, req.Timestamp, req.Nonce)
	if err != nil {
		log.Logger.Errorf("GenSendMsg error:%s", err.Error())
		return
	}

	ctx.Write(msg)
}

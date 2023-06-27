package controllers

import (
	ctx "big-genius/core/context"
	"big-genius/core/errors"
	"big-genius/core/utils"
	"big-genius/core/utils/wechat"
	"big-genius/internal/app/services/mq"
	"github.com/kataras/iris/v12/context"
	"strings"
	"time"
)

func Index(ctx ctx.Context) {
	word := "ChatGPT Robot is running"
	ctx.Reply(word, errors.GenErr(nil))
}

type AskReq struct {
	Question string `json:"question"`
}

func Ask(ctx ctx.Context) {
	var req AskReq

	req.Question = ctx.URLParam("question")

	if req.Question == "" {
		if err := ctx.ReadJSON(&req); err != nil {
			ctx.Reply(nil, errors.GenErr(err))
			return
		}
	}

	//s := services.NewOpenAIService()
	//answer, err := s.CreateChatCompletion(ctx, req.Question)
	//ctx.Reply(answer, errors.GenErr(err))
}

func MockWxWebhook(ctx *context.Context) {
	var req AskReq
	if err := ctx.ReadJSON(&req); err != nil {
		utils.Reply(ctx, nil, errors.GenErr(err))
		return
	}

	if strings.TrimSpace(req.Question) == "" {
		utils.Reply(ctx, nil, errors.GenErr(nil, errors.ErrNoBadRequest))
		return
	}

	msgCont := wechat.MsgContent{
		ToUsername:   "ww3e03cf4760de0bcd",
		FromUsername: "LiMingZe",
		CreateTime:   uint32(time.Now().Unix()),
		MsgType:      "text",
		Content:      req.Question,
		Msgid:        "",
		Agentid:      1000002,
	}

	mq.SendMQMsg(msgCont)

}

package controllers

import (
	ctx "big-genius/core/context"
	"big-genius/core/errors"
	"big-genius/internal/app/services"
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

	s := services.NewOpenAIService()
	answer, err := s.CreateChatCompletion(ctx, req.Question)
	ctx.Reply(answer, errors.GenErr(err))
}

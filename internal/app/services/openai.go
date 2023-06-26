package services

import (
	ctx "big-genius/core/context"
	"big-genius/core/log"
	"big-genius/internal/app/dao"
	"encoding/json"
)

type OpenAIService struct {
	OpenAIDAO *dao.OpenAIDAO
}

var theOpenAIService = new(OpenAIService)

func NewOpenAIService() *OpenAIService {
	if theOpenAIService.OpenAIDAO == nil {
		theOpenAIService.OpenAIDAO = dao.NewOpenAIDAO()
	}
	return theOpenAIService
}

func (s *OpenAIService) CreateChatCompletion(ctx ctx.Context, msg string) (string, error) {
	ccc, _ := json.Marshal(ctx)
	log.Logger.Error(string(ccc))
	resp, err := s.OpenAIDAO.CreateChatCompletion(ctx, msg)
	ddd, _ := json.Marshal(ctx)
	log.Logger.Error(string(ddd))
	if err != nil {
		return "", err
	}

	answer := ""

	if len(resp.Choices) > 0 {
		answer = resp.Choices[0].Message.Content
	}

	return answer, nil
}

package services

import (
	ctx "big-genius/core/context"
	"big-genius/internal/app/dao"
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
	resp, err := s.OpenAIDAO.CreateChatCompletion(ctx, msg)
	if err != nil {
		return "", err
	}

	answer := ""

	if len(resp.Choices) > 0 {
		answer = resp.Choices[0].Message.Content
	}

	return answer, nil
}

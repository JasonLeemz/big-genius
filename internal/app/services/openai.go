package services

import (
	"big-genius/core/config"
	"big-genius/internal/app/dao"
	"context"
	"time"
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

func (s *OpenAIService) CreateChatCompletion(msg string) (string, error) {
	timeout := config.GlobalConfig.OpenAI.TimeOut * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	resp, err := s.OpenAIDAO.CreateChatCompletion(ctx, msg)

	if err != nil {
		return err.Error(), err
	}

	answer := ""

	if len(resp.Choices) > 0 {
		answer = resp.Choices[0].Message.Content
	}

	return answer, nil
}

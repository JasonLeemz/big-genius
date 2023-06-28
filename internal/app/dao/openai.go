package dao

import (
	ctx "big-genius/core/context"
	"big-genius/core/log"
	ai "big-genius/internal/app/models/openai"
	"context"
	"github.com/sashabaranov/go-openai"
)

type OpenAI interface {
	CreateChatCompletion(ctx ctx.Context, msg string) (*openai.ChatCompletionResponse, error)
}

type OpenAIDAO struct {
	client *openai.Client
}

func (ai *OpenAIDAO) CreateChatCompletion(ctx context.Context, msg string) (*openai.ChatCompletionResponse, error) {
	resp, err := ai.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: msg,
				},
			},
		},
	)

	if err != nil {
		log.Logger.Errorf("ChatCompletion error: %v\n", err)
		return nil, err
	}
	// resp.Choices[0].Message.Content
	log.Logger.Infof("ChatCompletion resp: %v", resp)
	return &resp, err
}

func NewOpenAIDAO() *OpenAIDAO {
	return &OpenAIDAO{
		//client: ai.ChatGPT,
		client: ai.Azure,
	}
}

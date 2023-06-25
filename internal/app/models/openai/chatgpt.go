package openai

import (
	"big-genius/core/config"
	"github.com/sashabaranov/go-openai"
)

var AI *openai.Client

func Init() {
	cfg := config.GlobalConfig.OpenAI
	AI = openai.NewClient(cfg.Token)
}

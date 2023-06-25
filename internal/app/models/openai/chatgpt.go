package openai

import (
	"big-genius/core/config"
	"fmt"
	"github.com/sashabaranov/go-openai"
)

var AI *openai.Client

func ChatGPTInit() {
	cfg := config.GlobalConfig.OpenAI
	AI = openai.NewClient(cfg.Token)

	fmt.Println(AI)
}

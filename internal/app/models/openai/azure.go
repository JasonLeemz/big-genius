package openai

import (
	"big-genius/core/config"
	"github.com/sashabaranov/go-openai"
)

var Azure *openai.Client

func InitAzure() {
	conf := openai.DefaultAzureConfig(config.GlobalConfig.OpenAI.Azure.Token, config.GlobalConfig.OpenAI.Azure.BaseURL)
	conf.APIVersion = config.GlobalConfig.OpenAI.Azure.ApiVersion
	conf.AzureModelMapperFunc = modelMap
	Azure = openai.NewClientWithConfig(conf)
}

func modelMap(m string) string {
	return config.GlobalConfig.OpenAI.Azure.Deployments
}

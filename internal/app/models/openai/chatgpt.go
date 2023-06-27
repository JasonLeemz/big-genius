package openai

import (
	"big-genius/core/config"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"net/http"
	"net/url"
	"time"
)

var ChatGPT *openai.Client

func InitChatGPT() {
	// 创建一个自定义的 Transport
	host := fmt.Sprintf("%s://%s:%s",
		config.GlobalConfig.Proxy.Schema,
		config.GlobalConfig.Proxy.Host,
		config.GlobalConfig.Proxy.Port)
	proxyUrl, err := url.Parse(host)
	if err != nil {
		panic(err)
	}

	conf := openai.DefaultConfig(config.GlobalConfig.OpenAI.ChatGPT.Token)
	conf.HTTPClient = &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		},
		Timeout: time.Second * 200,
	}
	conf.BaseURL = config.GlobalConfig.OpenAI.ChatGPT.BaseURL

	ChatGPT = openai.NewClientWithConfig(conf)
}

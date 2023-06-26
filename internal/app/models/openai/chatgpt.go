package openai

import (
	"big-genius/core/config"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"net/http"
	"net/url"
	"time"
)

var AI *openai.Client

func Init() {
	// 创建一个自定义的 Transport
	host := fmt.Sprintf("%s://%s:%s",
		config.GlobalConfig.Proxy.Schema,
		config.GlobalConfig.Proxy.Host,
		config.GlobalConfig.Proxy.Port)
	proxyUrl, err := url.Parse(host)
	if err != nil {
		panic(err)
	}

	conf := openai.DefaultConfig(config.GlobalConfig.OpenAI.Token)
	conf.HTTPClient = &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		},
		Timeout: time.Second * 10,
	}
	conf.BaseURL = config.GlobalConfig.OpenAI.BaseURL

	AI = openai.NewClientWithConfig(conf)
}

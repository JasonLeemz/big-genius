package openai

import (
	"big-genius/core/config"
	"github.com/sashabaranov/go-openai"
)

var Azure *openai.Client

func InitAzure() {
	// 创建一个自定义的 Transport
	//host := fmt.Sprintf("%s://%s:%s",
	//	config.GlobalConfig.Proxy.Schema,
	//	config.GlobalConfig.Proxy.Host,
	//	config.GlobalConfig.Proxy.Port)
	//proxyUrl, err := url.Parse(host)
	//if err != nil {
	//	panic(err)
	//}

	conf := openai.DefaultAzureConfig(config.GlobalConfig.OpenAI.Azure.Token, config.GlobalConfig.OpenAI.Azure.BaseURL)
	//conf.HTTPClient = &http.Client{
	//	Transport: &http.Transport{
	//		Proxy: http.ProxyURL(proxyUrl),
	//	},
	//	Timeout: time.Second * 200,
	//}

	Azure = openai.NewClientWithConfig(conf)
}

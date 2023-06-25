package proxy

import (
	"big-genius/core/config"
	"fmt"
	"net/http"
	"net/url"
)

func Init() {
	// 设置代理地址
	host := fmt.Sprintf("%s://%s:%s",
		config.GlobalConfig.Proxy.Schema,
		config.GlobalConfig.Proxy.Host,
		config.GlobalConfig.Proxy.Port)
	fmt.Println(host)
	proxyUrl, err := url.Parse(host)
	if err != nil {
		panic(err)
	}
	// 设置全局代理
	http.DefaultTransport = &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}
}

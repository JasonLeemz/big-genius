package mq

import (
	"big-genius/core/log"
	"big-genius/core/utils/wechat"
	"big-genius/internal/app/services"
	ctx2 "context"
	"encoding/json"
	"sync"
	"time"
)

func SendWxMsg(msg []byte) {
	log.Logger.Info(string(msg))

	msgCont := wechat.MsgContent{}
	err := json.Unmarshal(msg, &msgCont)
	if err != nil {
		log.Logger.Errorf("MsgContent Unmarsha Error:%s", err.Error())
		return
	}

	wg := sync.WaitGroup{}
	wg.Add(2)
	// 请求chatgpt
	go func() {
		defer wg.Done()

		s := services.NewOpenAIService()

		startTime := time.Now()
		answer, err := s.CreateChatCompletion(msgCont.Content)
		elapsedTime := time.Since(startTime) / time.Second
		log.Logger.Infof("请求openai耗时:%s", elapsedTime)

		if err != nil {
			log.Logger.Errorf("请求openai超时:%s", err.Error())
			msgCont.Content = "请求openai超时"
		} else {
			log.Logger.Infof(">>>>>ANSWER<<<<<:%s", answer)
			msgCont.Content = answer
		}
	}()

	// 获取wxToken可能会请求网络
	accessToken := ""
	go func() {
		defer wg.Done()

		timeout := 5 * time.Second
		c, cancel := ctx2.WithTimeout(ctx2.Background(), timeout)
		defer cancel()

		accessToken = wechat.GetAccessToken(c)
	}()

	wg.Wait()
	wechat.SendMsg(msgCont.ToUsername, msgCont.Content, accessToken)

}

func SendMQMsg(msg wechat.MsgContent) {
	str, _ := json.Marshal(msg)
	log.Logger.Infof("SendMQMsg:%s", str)
	s := NewMQService()
	// ai ai.chatgpt
	s.ProduceMessage(Exchange, RoutingKeyChatgpt, str)
}
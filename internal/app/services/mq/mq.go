package mq

import (
	"big-genius/core/log"
	"big-genius/internal/app/dao"
	"github.com/streadway/amqp"
)

type MQService struct {
	MQDAO *dao.MQDAO
}

var (
	Exchange = "ai"

	RoutingKeyChatgpt = "ai.chatgpt"
	RoutingKeyWXYY    = "ai.wxyy"

	QueneNameChatgpt = "chatgpt"
	QueneNameWXYY    = "wxyy" // 文心一言
)

// ProduceMessage exchange: ai, queneName: ai.*
func (s *MQService) ProduceMessage(exchange, routingKey string, message []byte) {
	ch, err := s.MQDAO.Conn.Channel()
	if err != nil {
		log.Logger.Error(err)
		return
	}

	err = ch.Publish(
		exchange,   // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		})
	if err != nil {
		log.Logger.Errorf("Failed to publish a message: %s", err.Error())
	} else {
		log.Logger.Infof("%s ok", routingKey)
	}
}

func (s *MQService) ConsumeMessagesForChatgpt() {
	queueName := QueneNameChatgpt
	ch, err := s.MQDAO.Conn.Channel()
	if err != nil {
		log.Logger.Panic(err)
		panic(err)
	}

	msgs, err := ch.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		log.Logger.Error("Failed to register a consumer")
		log.Logger.Panic(err)
		panic(err)
	}

	//forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Logger.Infof("Received a message: %s", d.Body)
			SendWxMsg(d.Body)
		}
	}()

	log.Logger.Info("Waiting for messages...")
	//<-forever
}

var theMQService = new(MQService)

func NewMQService() *MQService {
	if theMQService.MQDAO == nil {
		theMQService.MQDAO = dao.NewMQDAO()
	}
	return theMQService
}

func RegisterTrigger() {
	NewMQService().ConsumeMessagesForChatgpt()
}

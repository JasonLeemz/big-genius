package dao

import (
	"big-genius/core/log"
	"big-genius/internal/app/models/mq"
	"github.com/streadway/amqp"
)

type MQ interface {
	ProduceMessage(ch *amqp.Channel, queueName, message string)
	ConsumeMessages(ch *amqp.Channel, queueName string)
}

type MQDAO struct {
	Conn *amqp.Connection
}

func (dao *MQDAO) ProduceMessage(ch *amqp.Channel, queueName, message string) error {
	return ch.Publish(
		"",        // exchange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

func (dao *MQDAO) ConsumeMessages(ch *amqp.Channel, queueName string) {
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
		log.Logger.Errorf("Failed to register a consumer:%s", err.Error())
		panic(err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Logger.Infof("Received a message: %s", d.Body)

		}
	}()

	log.Logger.Info("Waiting for messages...")
	<-forever
}

func NewMQDAO() *MQDAO {
	return &MQDAO{
		Conn: mq.Conn,
	}
}

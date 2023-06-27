package mq

import (
	"big-genius/core/config"
	"big-genius/core/log"
	"fmt"
	"github.com/streadway/amqp"
)

var Conn *amqp.Connection

func failOnError(err error, msg string) {
	if err != nil {
		log.Logger.Errorf("%s: %s", msg, err)
	}
}

func Init() {
	// 连接 RabbitMQ
	url := fmt.Sprintf("%s://%s:%s@%s:%s",
		config.GlobalConfig.MQ.Schema,
		config.GlobalConfig.MQ.Username,
		config.GlobalConfig.MQ.Password,
		config.GlobalConfig.MQ.Host,
		config.GlobalConfig.MQ.Port)
	var err error
	Conn, err = amqp.Dial(url)
	failOnError(err, "Failed to connect to RabbitMQ")
	//defer Conn.Close()
}

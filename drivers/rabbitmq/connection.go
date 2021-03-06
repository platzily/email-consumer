package rabbitmq

import (
	"github.com/platzily/email-consumer/config"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func NewConnection() (conn *amqp.Connection, err error) {

	env := config.ReadRabbitMQConfig()
	conn, err = amqp.Dial(env.URL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	defer conn.Close()

	log.Infof("Rabbit Connection Success to %s", env.URL)

	return conn, nil
}

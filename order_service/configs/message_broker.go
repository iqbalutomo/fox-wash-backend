package config

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

func ConnectMessageBroker() (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		log.Fatalf("failed to connect rabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("failed to connect channel: %v", err)
	}

	return conn, ch
}

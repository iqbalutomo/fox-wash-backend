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

func InitMessageBrokerQueue(ch *amqp.Channel, nameType string) amqp.Queue {
	q, err := ch.QueueDeclare(
		nameType, // name
		false,    // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		log.Fatalf("failed to initialize queue: %v", err)
	}

	return q
}

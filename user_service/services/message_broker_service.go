package services

import (
	"log"

	"github.com/streadway/amqp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MessageBroker interface {
	PublishMessage(message []byte) error
}

type RabbitMQ struct {
	ch *amqp.Channel
}

func NewMessageBroker(ch *amqp.Channel) MessageBroker {
	return &RabbitMQ{ch}
}

func (r *RabbitMQ) PublishMessage(message []byte) error {
	q, err := r.ch.QueueDeclare(
		"email_verification", // name
		false,                // durable
		false,                // delete when unused
		false,                // exclusive
		false,                // no-wait
		nil,                  // arguments
	)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	if err := r.ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		},
	); err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	log.Printf(" [x] sent %s\n", message)

	return nil
}

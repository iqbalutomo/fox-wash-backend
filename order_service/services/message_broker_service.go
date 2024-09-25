package services

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MessageBroker interface {
	PublishMessageOrder(message map[string]interface{}) error
}

type RabbitMQ struct {
	ch *amqp.Channel
}

func NewMessageBroker(ch *amqp.Channel) MessageBroker {
	return &RabbitMQ{ch}
}

func (r *RabbitMQ) PublishMessageOrder(message map[string]interface{}) error {
	return r.PublishMessage("email_order_user", message)
}

func (r *RabbitMQ) PublishMessage(queueName string, message map[string]interface{}) error {
	q, err := r.ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	if err := r.ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	); err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	log.Printf(" [x] sent %s to %s\n", message, queueName)

	return nil
}

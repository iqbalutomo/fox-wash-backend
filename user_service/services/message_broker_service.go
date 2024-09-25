package services

import (
	"log"

	"github.com/streadway/amqp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MessageBroker interface {
	PublishMessageVerification(message []byte) error
	PublishMessagePaymentSuccess(message []byte) error
}

type RabbitMQ struct {
	ch *amqp.Channel
}

func NewMessageBroker(ch *amqp.Channel) MessageBroker {
	return &RabbitMQ{ch}
}

func (r *RabbitMQ) PublishMessageVerification(message []byte) error {
	return r.PublishMessage("email_verification", message)
}

func (r *RabbitMQ) PublishMessagePaymentSuccess(message []byte) error {
	return r.PublishMessage("email_payment_success", message)
}

func (r *RabbitMQ) PublishMessage(queueName string, message []byte) error {
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

	log.Printf(" [x] sent %s to %s\n", message, queueName)

	return nil
}

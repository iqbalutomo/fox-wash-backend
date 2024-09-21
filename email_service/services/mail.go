package services

import (
	"email_service/helpers"
	"log"

	"github.com/streadway/amqp"
)

type Mail interface {
	SendEmailVerification(amqp.Queue)
}

type MailService struct {
	channel *amqp.Channel
}

func NewMailService(channel *amqp.Channel) Mail {
	return &MailService{channel}
}

func (m *MailService) SendEmailVerification(q amqp.Queue) {
	msgs, err := m.channel.Consume(
		q.Name, //queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("failed to consuming message: %v", err)
	}

	for d := range msgs {
		log.Printf("new message: %s", d.Body)

		userData := helpers.AssertJsonToStruct(d.Body)
		if err := helpers.SendEmailVerification(userData); err != nil {
			log.Fatal(err)
		}
	}
}

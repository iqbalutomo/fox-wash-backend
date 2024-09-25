package services

import (
	"email_service/helpers"
	"email_service/models"
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

type Mail interface {
	SendEmailVerification(q amqp.Queue)
	SendEmailOrder(q amqp.Queue)
	SendEmailPaymentSuccess(q amqp.Queue)
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
		log.Fatalf("failed to consume verification message: %v", err)
	}

	for d := range msgs {
		log.Printf("new verification message: %s", d.Body)

		userData := helpers.AssertJsonToStruct(d.Body)
		if err := helpers.SendEmailVerification(userData); err != nil {
			log.Fatal(err)
		}
	}
}

func (m *MailService) SendEmailOrder(q amqp.Queue) {
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
		log.Fatalf("failed to consume order message: %v", err)
	}

	for d := range msgs {
		log.Printf("new order message: %s", d.Body)

		var orderData models.Order

		if err := json.Unmarshal(d.Body, &orderData); err != nil {
			log.Fatalf("failed to unmarshaling data: %v", err)
		}

		if err := helpers.SendEmailOrder(orderData); err != nil {
			log.Fatal(err)
		}
	}
}

func (m *MailService) SendEmailPaymentSuccess(q amqp.Queue) {
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
		log.Fatalf("failed to consume payment success message: %v", err)
	}

	for d := range msgs {
		log.Printf("new order message: %s", d.Body)

		var paymentSuccessData models.PaymentSuccess

		if err := json.Unmarshal(d.Body, &paymentSuccessData); err != nil {
			log.Fatalf("failed to unmarshaling data: %v", err)
		}

		if err := helpers.SendPaymentSuccess(paymentSuccessData); err != nil {
			log.Fatal(err)
		}
	}
}

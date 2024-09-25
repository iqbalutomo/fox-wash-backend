package main

import (
	config "email_service/configs"
	"email_service/services"
	"log"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	conn, mbChan := config.ConnectMessageBroker()
	defer conn.Close()

	qVerification := config.InitMessageBrokerQueue(mbChan, "email_verification")
	qOrder := config.InitMessageBrokerQueue(mbChan, "email_order_user")
	mailService := services.NewMailService(mbChan)

	go mailService.SendEmailVerification(qVerification)
	go mailService.SendEmailOrder(qOrder)

	forever := make(chan bool)
	log.Printf(" [*] waiting for messages. to exit press CTRL+C\n")
	<-forever
}

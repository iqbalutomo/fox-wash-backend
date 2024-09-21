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

	q := config.InitMessageBrokerQueue(mbChan)
	mailService := services.NewMailService(mbChan)

	go mailService.SendEmailVerification(q)

	forever := make(chan bool)
	log.Printf(" [*] waiting for messages. to exit press CTRL+C\n")
	<-forever
}

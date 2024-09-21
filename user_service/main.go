package main

import (
	"log"
	config "user_service/configs"
	"user_service/controllers"
	"user_service/repository"
	"user_service/services"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	conn, mbChan := config.ConnectMessageBroker()
	defer conn.Close()

	messageBrokerService := services.NewMessageBroker(mbChan)
	userRepo := repository.NewUserRepository(db)
	userController := controllers.NewUserController(userRepo, messageBrokerService)

	config.ListenAndServeGrpc(*userController)
}

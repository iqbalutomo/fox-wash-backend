package main

import (
	"log"
	config "user_service/configs"
	"user_service/controllers"
	"user_service/repository"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	userRepo := repository.NewUserRepository(db)
	userController := controllers.NewUserController(userRepo)

	config.ListenAndServeGrpc(*userController)
}

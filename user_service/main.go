package main

import (
	"log"
	config "user_service/configs"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	_ = db

	config.ListenAndServeGrpc()
}

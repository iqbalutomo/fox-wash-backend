package main

import (
	"log"
	config "wash_station_service/configs"
	"wash_station_service/controllers"
	"wash_station_service/repository"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	washStationRepo := repository.NewWashStationRepository(db)
	washStationController := controllers.NewWashStationController(washStationRepo)

	config.ListenAndServeGrpc(*washStationController)
}

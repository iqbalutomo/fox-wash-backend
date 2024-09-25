package main

import (
	config "order_service/configs"
	"order_service/controllers"
	"order_service/repository"
	"order_service/services"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	db := config.ConnectDB().Database(os.Getenv("MONGO_DBNAME"))
	orderCollection := db.Collection(os.Getenv("MONGO_COLLECTION"))

	conn, userServiceClient := config.InitUserServiceClient()
	defer conn.Close()

	conn, washstationClient := config.InitWashStationServiceClient()
	defer conn.Close()

	paymentService := services.NewPaymentService(os.Getenv("XENDIT_API_KEY"))
	orderRepo := repository.NewOrderRepository(orderCollection)
	orderController := controllers.NewOrderController(orderRepo, userServiceClient, washstationClient, paymentService)

	config.ListenAndServeGrpc(*orderController)
}

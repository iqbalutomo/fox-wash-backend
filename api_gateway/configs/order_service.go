package config

import (
	"api_gateway/pb/orderpb"
	"log"
	"os"

	"google.golang.org/grpc"
)

func InitOrderServiceClient() (*grpc.ClientConn, orderpb.OrderServiceClient) {
	conn, err := grpc.NewClient(os.Getenv("ORDER_SERVICE_URI"), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return conn, orderpb.NewOrderServiceClient(conn)
}

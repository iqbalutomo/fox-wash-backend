package config

import (
	"log"
	"order_service/pb/userpb"
	"os"

	"google.golang.org/grpc"
)

func InitUserServiceClient() (*grpc.ClientConn, userpb.UserClient) {
	conn, err := grpc.NewClient(os.Getenv("USER_SERVICE_URI"), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return conn, userpb.NewUserClient(conn)
}

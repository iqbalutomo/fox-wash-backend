package config

import (
	pb "api_gateway/pb/userpb"
	"log"
	"os"

	"google.golang.org/grpc"
)

func InitUserServiceClient() (*grpc.ClientConn, pb.UserClient) {
	conn, err := grpc.NewClient(os.Getenv("USER_SERVICE_URI"), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return conn, pb.NewUserClient(conn)
}

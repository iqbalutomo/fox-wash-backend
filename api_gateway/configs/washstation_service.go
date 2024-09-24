package config

import (
	pb "api_gateway/pb/washstationpb"
	"log"
	"os"

	"google.golang.org/grpc"
)

func InitWashStationServiceClient() (*grpc.ClientConn, pb.WashStationClient) {
	conn, err := grpc.NewClient(os.Getenv("WASHSTATION_SERVICE_URI"), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return conn, pb.NewWashStationClient(conn)
}

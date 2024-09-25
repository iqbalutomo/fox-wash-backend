package config

import (
	"log"
	"order_service/pb/washstationpb"

	"os"

	"google.golang.org/grpc"
)

func InitWashStationServiceClient() (*grpc.ClientConn, washstationpb.WashStationClient) {
	conn, err := grpc.NewClient(os.Getenv("WASHSTATION_SERVICE_URI"), grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return conn, washstationpb.NewWashStationClient(conn)
}

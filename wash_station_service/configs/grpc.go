package config

import (
	"log"
	"net"
	"os"
	"wash_station_service/controllers"
	"wash_station_service/middlewares"
	"wash_station_service/pb"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"google.golang.org/grpc"
)

func ListenAndServeGrpc(controller controllers.Server) {
	port := os.Getenv("PORT")
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpc_auth.UnaryServerInterceptor(middlewares.JWTAuth),
		),
	)

	pb.RegisterWashStationServer(grpcServer, &controller)

	log.Println("gRPC server is running on port:", port)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to gRPC server: %v", err)
	}
}

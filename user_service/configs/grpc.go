package config

import (
	"log"
	"net"
	"os"
	"user_service/controllers"
	"user_service/pb"

	"google.golang.org/grpc"
)

func ListenAndServeGrpc(controller controllers.Server) {
	port := os.Getenv("PORT")
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterUserServer(grpcServer, &controller)

	log.Println("gRPC server is running on port:", port)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to gRPC server: %v", err)
	}
}

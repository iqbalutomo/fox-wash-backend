package config

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func ListenAndServeGrpc() {
	port := os.Getenv("PORT")
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	log.Println("gRPC server is running on port:", port)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to gRPC server: %v", err)
	}
}

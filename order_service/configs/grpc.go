package config

import (
	"log"
	"net"
	"order_service/controllers"
	"order_service/middlewares"
	"order_service/pb/orderpb"
	"os"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"

	"google.golang.org/grpc"
)

func ListenAndServeGrpc(controller controllers.OrderController) {
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

	orderpb.RegisterOrderServiceServer(grpcServer, &controller)

	log.Println("gRPC server is running on port:", port)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to gRPC server: %v", err)
	}
}

package main

import (
	"fmt"
	"log"
	"net"

	"github.com/dixneuf19/go-min-docker/config"
	helloword "github.com/dixneuf19/go-min-docker/grpc"

	"google.golang.org/grpc"
)

func main() {
	config.Init()
	fmt.Println("App initiated")

	// create a listener on TCP port 50051
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// create a server instance
	s := helloword.Server{}
	// create a gRPC server object
	grpcServer := grpc.NewServer()
	// attach the Ping service to the server
	helloword.RegisterGreeterServer(grpcServer, &s)
	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

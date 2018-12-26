package helloworld

import (
	"fmt"
	"log"

	"golang.org/x/net/context"
)

// Server represents the gRPC server
type Server struct {
}

// SayHello generates response to a Ping request
func (s *Server) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	log.Printf("Receive message %s", in.GetName())
	return &HelloReply{Message: fmt.Sprintf("Hello %s !", in.GetName())}, nil
}

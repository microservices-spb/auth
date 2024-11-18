package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/microservices-spb/auth/internal/rpc"
	"github.com/microservices-spb/auth/pkg/auth"
)

func main() {
	service := rpc.NewHandler()

	s := grpc.NewServer()

	auth.RegisterAuthServiceServer(s, service)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", "3112"))
	if err != nil {
		log.Printf("Cannot listen port:")
	}

	s.Serve(lis)
}

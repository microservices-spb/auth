package main

import (
	"auth-service/internal/rpc"
	"auth-service/pkg/auth"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
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

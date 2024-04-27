package main

import (
	"log"
	"net"
	"unary/proto/service"

	"google.golang.org/grpc"
)

func main() {

	listen, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	service.RegisterUnaryServiceServer(server, Server{})

	err = server.Serve(listen)
	if err != nil {
		log.Fatal(err)
	}
}

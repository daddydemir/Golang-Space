package main

import (
	"log"
	"net"
	"server-stream/proto/service"

	"google.golang.org/grpc"
)

func main() {

	listen, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	service.RegisterServerServiceServer(server, Server{})

	err = server.Serve(listen)
	if err != nil {
		log.Fatal(err)
	}
}

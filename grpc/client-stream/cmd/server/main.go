package main

import (
	"client-stream/proto/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	listen, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	service.RegisterClientServiceServer(server, Server{})

	err = server.Serve(listen)
	if err != nil {
		log.Fatal(err)
	}

}

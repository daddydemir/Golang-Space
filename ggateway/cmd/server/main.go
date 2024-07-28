package main

import (
	"ggateway/internal"
	users "ggateway/protogen/golang/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	const addr = ":5001"

	server := grpc.NewServer()

	users.RegisterUsersServer(server, &internal.UserService{})

	reflection.Register(server)

	log.Printf("server started at port: %v \n", addr)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}

	if err = server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

package main

import (
	"context"
	"grpc/proto"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	proto.RegisterUserServiceServer(server, UserServer{})

	err = server.Serve(listener)
	if err != nil {
		log.Println(err)
	}

}

type UserServer struct{}

// GetUserStream implements proto.UserServiceServer.
func (UserServer) GetUserStream(stream proto.UserService_GetUserStreamServer) error {

	for {
		ur, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		key := ur.Name

		if err = stream.Send(&proto.UserReply{Name: key}); err != nil {
			return err
		}
	}
}

// GetUserStreamReply implements proto.UserServiceServer.
func (UserServer) GetUserStreamReply(request *proto.UserRequest, stream proto.UserService_GetUserStreamReplyServer) error {

	for {
		stream.Send(&proto.UserReply{Name: "Hello " + request.Name})
		return nil
	}
}

// GetUserStreamRequest implements proto.UserServiceServer.
func (UserServer) GetUserStreamRequest(stream proto.UserService_GetUserStreamRequestServer) error {

	names := make([]string, 2)

	for {
		ur, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&proto.UserReply{Name: "Hello " + names[3]})
		}
		if err != nil {
			return err
		}

		names = append(names, ur.Name)
	}
}

// GetUser implements proto.UserServiceServer.
func (UserServer) GetUser(ctx context.Context, request *proto.UserRequest) (*proto.UserReply, error) {

	log.Println("getUser request:", request.Name)

	return &proto.UserReply{Name: request.Name}, nil
}

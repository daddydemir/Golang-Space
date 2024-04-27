package main

import (
	"bidirectional-stream/proto/message"
	"bidirectional-stream/proto/service"
	"context"
	"io"
	"log"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:5001", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := service.NewBidirectionalStreamClient(conn)

	chat, err := client.StartChat(context.Background())
	if err != nil {
		log.Printf("Client ERROR: %v", err)
	}

	requests := []*message.Request{
		{Message: "Hello World!"},
		{Message: "Hello Server."},
	}

	go func() {
		for _, req := range requests {
			if err = chat.Send(req); err != nil {
				log.Fatalf("Error sending request: %v", err)
			}
		}
		chat.CloseSend()
	}()

	for {
		response, err := chat.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Client ERROR: %v", err)
		}
		log.Printf("CLIENT: %v \n", response.Message)
	}

}

package main

import (
	"client-stream/proto/message"
	"client-stream/proto/service"
	"context"
	"log"

	"google.golang.org/grpc"
)

var file = [][]byte{
	[]byte("file part-1"),
	[]byte("file part-2"),
	[]byte("file part-3"),
	[]byte("file part-4"),
	[]byte("file part-5"),
}

func main() {
	conn, err := grpc.Dial("localhost:5001", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := service.NewClientServiceClient(conn)
	ctx := context.Background()
	stream, err := client.UploadFile(ctx)
	if err != nil {
		log.Printf("Client ERROR: %v", err)
	}

	for _, data := range file {
		err := stream.Send(&message.Request{ChunkData: data})
		if err != nil {
			log.Printf("Client ERROR: %v", err)
		}
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("Client ERROR: %v", err)
	}
	log.Printf("Response: %v \n", response.Message)
}

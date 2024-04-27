package main

import (
	"context"
	"io"
	"log"
	"server-stream/proto/message"
	"server-stream/proto/service"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:5001", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := service.NewServerServiceClient(conn)
	ctx := context.Background()

	stream, err := client.GetVideoStream(ctx, &message.Request{VideoId: "1337"})
	if err != nil {
		log.Printf("Client ERROR: %v \n", err)
	}

	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Recv error: %v", err)
		}
		log.Printf("Received video chunk: %v", string(chunk.ChunkData))
	}
}

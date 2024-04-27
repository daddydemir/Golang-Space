package main

import (
	"context"
	"log"
	"unary/proto/message"
	"unary/proto/service"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:5001", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := service.NewUnaryServiceClient(conn)
	ctx := context.Background()

	for i := 0; i < 3; i++ {
		product, err := client.GetProduct(ctx, &message.Request{Values: int32(i)})
		if err != nil {
			log.Printf("Client ERROR: %v \n", err)
		}
		log.Println(product)
	}

}

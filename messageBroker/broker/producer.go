package broker

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func StartProducer() {

	Writer := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "SecondTopic",
		Balancer: &kafka.LeastBytes{},
	}

	fmt.Printf("send topic: %v \n", Writer.Topic)
	index := 0

	for index < 15 {
		index++

		message := kafka.Message{
			Value: []byte(fmt.Sprintf("Message-%d", index)),
		}

		err := Writer.WriteMessages(context.Background(), message)
		if err != nil {
			log.Fatalf("WriteMessage occurred error:%v", err)
		}
		fmt.Printf("WRITER: %v  partition: %v \n", string(message.Value), message.Partition)

	}
}

package broker

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func StartConsumer() {

	Reader := &kafka.ReaderConfig{
		Brokers:     []string{"localhost:9092"},
		Topic:       "SecondTopic",
		GroupID:     "group-1",
		StartOffset: kafka.LastOffset,
	}

	r := kafka.NewReader(*Reader)

	messageChannel := make(chan *kafka.Message)

	fmt.Printf("listen on topic: %v \n", r.Config().Topic)

	go func() {
		for {
			m, err := r.ReadMessage(context.Background())
			if err != nil {
				log.Fatalf("ReadMessage has occurred error:%v", err)
			}
			messageChannel <- &m
		}
	}()

	for msg := range messageChannel {
		fmt.Printf("READER message:%v offset:%v partition:%v \n", string(msg.Value), msg.Offset, msg.Partition)
	}

}

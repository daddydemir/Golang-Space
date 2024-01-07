package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func createTopic() {
	c, err := kafka.Dial("tcp", "localhost:9092")
	if err != nil {
		log.Print(err)
	}

	err = c.CreateTopics([]kafka.TopicConfig{{
		Topic:             "gracefulShutdown",
		NumPartitions:     3,
		ReplicationFactor: 1,
	}}...)
	if err != nil {
		log.Print(err)
	}
}

func Produce() {

	writer := kafka.Writer{
		Addr:  kafka.TCP("localhost:9092"),
		Topic: "gracefulShutdown",
	}

	index := 1
	for {
		msg := kafka.Message{
			Value: []byte(fmt.Sprintf("Message Number : %d", index)),
		}

		writer.WriteMessages(context.Background(), msg)
		index++
	}
}

func consume() {

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{"localhost:9092"},
		Topic:       "gracefulShutdown",
		GroupID:     "group-1",
		StartOffset: kafka.LastOffset,
	})

	defer reader.Close()

	for {
		m, err := reader.FetchMessage(context.Background())
		if err != nil {
			log.Println("read error: ", err)
		}

		log.Println("message: ", m)
		time.Sleep(2 * time.Second)

		err = reader.CommitMessages(context.Background(), m)
		if err != nil {
			log.Println("commit error: ", err)
		} else {
			log.Println("commit successfully : ", m)
		}
	}

}

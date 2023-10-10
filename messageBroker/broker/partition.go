package broker

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func GetDetails() {
	broker := "localhost:9092"
	topic := "SecondTopic"

	conn, err := kafka.DialContext(context.Background(), "tcp", broker)
	if err != nil {
		log.Fatalf("Kafka broker'a bağlanma hatası: %v", err)
	}
	defer conn.Close()

	metadata, err := conn.ReadPartitions(topic)
	if err != nil {
		log.Fatalf("ReadPartitions err: %v", err)
	}

	partitionCount := len(metadata)

	fmt.Printf("'%s' konusu %d partition'a sahiptir.\n", topic, partitionCount)
}

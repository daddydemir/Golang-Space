package broker

import (
	"log"
	"net"
	"strconv"

	"github.com/segmentio/kafka-go"
)

func CreateTopics() {

	conn, err := kafka.Dial("tcp", "localhost:9092")
	if err != nil {
		log.Fatalf("kafka connection creation has occurred error:%v", err)
	}
	defer conn.Close()

	controller, _ := conn.Controller()

	var controllerConn *kafka.Conn
	controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		panic(err.Error())
	}

	defer controllerConn.Close()

	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             "ThirdTopic",
			NumPartitions:     3,
			ReplicationFactor: 1,
		},
	}

	err = controllerConn.CreateTopics(topicConfigs...)
	if err != nil {
		panic(err.Error())
	}

}

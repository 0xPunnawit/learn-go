package main

import (
	"fmt"

	"github.com/IBM/sarama"
)

func main() {

	server := []string{"localhost:9092"}

	producer, err := sarama.NewSyncProducer(server, nil)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	msg := sarama.ProducerMessage{
		Topic: "inghello",
		Value: sarama.StringEncoder("hello world"),
	}

	p, o, err := producer.SendMessage(&msg)
	if err != nil {
		panic(err)
	}
	fmt.Println("partition=%v, offset=%v", p, o)

}

package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
)

func main() {
	conf := kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "github-events",
		GroupID: "g1",
	}

	reader := kafka.NewReader(conf)
	fmt.Println("Started listening...")
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("Error occurred", err)
			continue
		}

		fmt.Println("Message is:", string(msg.Value))
	}
}

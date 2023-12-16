package utils

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
)

func ListenToUpcomingEvents(ch chan kafka.Message) {
	conf := kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "github-events",
	}

	reader := kafka.NewReader(conf)
	fmt.Println("Started listening...")
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("Error occurred", err)
			continue
		}

		fmt.Println("Received Message:", string(msg.Value))
		ch <- msg
	}
}

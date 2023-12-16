package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
)

func main() {
	//ch := make(chan kafka.Message)
	//var wg sync.WaitGroup
	//
	//wg.Add(1)
	//go func() {
	//	fmt.Println("Hree")
	//	utils.ListenToUpcomingEvents(ch)
	//	wg.Done()
	//}()
	//
	//wg.Add(1)
	//go func() {
	//	utils.ProcessEvents(ch)
	//	wg.Done()
	//}()
	//
	//wg.Wait()

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
	}
}

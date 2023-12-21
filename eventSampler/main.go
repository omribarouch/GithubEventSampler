package main

import (
	"GithubEventHandler/utils"
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"time"
)

func main() {
	kafkaWriter := kafka.Writer{
		Addr:  kafka.TCP("localhost:9092"),
		Topic: "github-events",
	}

	for {
		fmt.Println("Fetching events...")
		githubEvents := utils.SendGithubRequest("GET", "https://api.github.com/events")
		err := kafkaWriter.WriteMessages(context.Background(), kafka.Message{
			Value: githubEvents,
		})
		if err != nil {
			fmt.Println("Error occurred while writing events to kafka:", err)
		}

		time.Sleep(5 * time.Second)
	}
}

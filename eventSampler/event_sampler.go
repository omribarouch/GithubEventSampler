package main

import (
	"context"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/google/go-github/v38/github"
	"golang.org/x/oauth2"
)

func main() {
	// Set up a context for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// GitHub API token, Kafka broker, and topic
	token := "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"
	broker := "localhost:9093"
	topic := "github-events"

	// Kafka producer configuration
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %s", err)
	}
	defer producer.Close()

	// GitHub API client setup
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// Periodically fetch and send GitHub events to Kafka
	ticker := time.NewTicker(2 * time.Minute) // Adjust the interval as needed

	for {
		select {
		case <-ctx.Done():
			log.Println("Event Sampler shutting down.")
			return
		case <-ticker.C:
			events, _, err := client.Activity.ListEvents(ctx, nil)
			if err != nil {
				log.Printf("Error fetching GitHub events: %s", err)
				continue
			}

			for _, event := range events {
				// Produce the message to the specified topic
				err := producer.Produce(&kafka.Message{
					TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
					Value:          event.GetRawPayload(),
				}, nil)

				if err != nil {
					log.Printf("Failed to produce message: %v", err)
				}
			}
		}
	}
}

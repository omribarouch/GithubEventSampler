package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"io"
	"net/http"
	"time"
)

func main() {
	kafkaWriter := kafka.Writer{
		Addr:  kafka.TCP("localhost:9092"),
		Topic: "github-events",
	}

	for {
		fmt.Println("Fetching events...")
		githubEvents := FetchGithubEvents()
		err := kafkaWriter.WriteMessages(context.Background(), kafka.Message{
			Value: githubEvents,
		})
		if err != nil {
			fmt.Println("Error occurred while writing events to kafka:", err)
		}

		fmt.Println("Now sleeping...")
		time.Sleep(5 * time.Second)
	}
}

func FetchGithubEvents() []byte {
	req, err := http.NewRequest("GET", "https://api.github.com/events", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
	}
	req.Header.Set("Authorization", "Bearer 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")
	req.Header.Set("Accept", "application/vnd.github+json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error performing request:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: Unexpected status code", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
	}

	return body
}

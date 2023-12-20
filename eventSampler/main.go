package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	kafkaWriter := kafka.Writer{
		Addr:  kafka.TCP(os.Getenv("KAFKA_BROKER")),
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

		time.Sleep(5 * time.Second)
	}
}

func FetchGithubEvents() []byte {
	req, err := http.NewRequest("GET", "https://api.github.com/events", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
	}
	req.Header.Set("Authorization", "Bearer "+os.Getenv("GITHUB_TOKEN"))
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

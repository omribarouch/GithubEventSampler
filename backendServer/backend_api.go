package backendServer

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func BackendAPI(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	switch vars["action"] {
	case "listEvents":
		listEvents(w, r)
	case "listRecentActors":
		listRecentActors(w, r)
	case "listRecentRepositories":
		listRecentRepositories(w, r)
	default:
		http.Error(w, "Invalid action", http.StatusBadRequest)
	}
}

func listEvents(w http.ResponseWriter, r *http.Request) {
	// Implement logic to retrieve and return all collected events
	// from the database or cache
	// ...

	// Example response
	response := []string{"Event1", "Event2", "Event3"}
	sendJSONResponse(w, response)
}

func listRecentActors(w http.ResponseWriter, r *http.Request) {
	// Implement logic to retrieve and return the 20 most recent actors
	// from the database or cache
	// ...

	// Example response
	response := []string{"Actor1", "Actor2", "Actor3"}
	sendJSONResponse(w, response)
}

func listRecentRepositories(w http.ResponseWriter, r *http.Request) {
	// Implement logic to retrieve and return the 20 most recent repositories
	// from the database or cache
	// ...

	// Example response
	response := []string{"Repo1", "Repo2", "Repo3"}
	sendJSONResponse(w, response)
}

func sendJSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("Error encoding JSON response: %s", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func main() {
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

		// Kafka broker address
		broker := "localhost:9093"

		// Kafka consumer configuration
		consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
			"bootstrap.servers":  broker,
			"group.id":           "test-group",
			"auto.offset.reset":  "earliest", // Adjust based on your requirements
			"enable.auto.commit": "false",    // Disable auto-commit offsets
		})
		if err != nil {
			log.Fatalf("Failed to create Kafka consumer: %s", err)
		}
		defer consumer.Close()

		// Subscribe to the Kafka topic
		topic := "github-events"
		err = consumer.SubscribeTopics([]string{topic}, nil)
		if err != nil {
			log.Fatalf("Error subscribing to topic: %s", err)
		}

		// Consume messages from Kafka
		for {
			select {
			case sig := <-sigChan:
				log.Printf("Received signal %s. Shutting down consumer...", sig)
				return
			default:
				msg, err := consumer.ReadMessage(-1)
				if err != nil {
					log.Printf("Error reading message: %v", err)
					continue
				}

				// Process the message (e.g., print it)
				fmt.Printf("Received message: %s\n", string(msg.Value))

				// Acknowledge the message manually since auto-commit is disabled
				consumer.CommitMessage(msg)
			}
		}
	}()

	router := mux.NewRouter()
	router.HandleFunc("/api/{action}", BackendAPI).Methods("GET")

	log.Println("Backend API server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

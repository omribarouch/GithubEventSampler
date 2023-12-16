package main

import (
	"GithubEventHandler/database"
	"GithubEventHandler/eventProcessor/utils"
	"github.com/segmentio/kafka-go"
	"sync"
)

func main() {
	database.ConnectDB()

	ch := make(chan kafka.Message, 1000)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		utils.ListenToUpcomingEvents(ch)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		utils.ProcessEvents(ch)
	}()

	wg.Wait()
}

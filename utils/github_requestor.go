package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func SendGithubRequest(requestType string, requestUrl string) []byte {
	req, err := http.NewRequest(requestType, requestUrl, nil)
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

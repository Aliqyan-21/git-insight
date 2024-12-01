package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

// fetchUserEvents function fetches the github events for the username provided to us
func fetchUserEvents(username string) (string, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	req, err := http.Get(url)

	if req.StatusCode != http.StatusOK {
		return "", errors.New(fmt.Sprintf("failed to fetch events %s", req.Status))
	}
	if err != nil {
		return "", err
	}

	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: git-insight <username>")
		return
	}

	username := os.Args[1]
	events, err := fetchUserEvents(username)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Events fetched successfully:")
	fmt.Println(events)
}

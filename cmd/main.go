package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Event struct {
	Type string `json:"type"`
	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`
}

// fetchUserEvents function fetches the github events for the username provided to us
func fetchUserEvents(username string) ([]Event, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	req, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer req.Body.Close()

	if req.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("failed to fetch events %s", req.Status))
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	var events []Event
	err = json.Unmarshal(body, &events)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error parsing json: %s", err))
	}

	return events, nil
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

	fmt.Printf("Events fetched successfully: \n\n")
	for _, event := range events {
		fmt.Printf("%s: %s\n", event.Type, event.Repo.Name)
	}
}

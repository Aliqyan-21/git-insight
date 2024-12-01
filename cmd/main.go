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

func checkEvents(typeOfEvent string) bool {
	// definition of relevant events for us
	relevantEvents := map[string]bool{
		"PushEvent":         true,
		"PullRequestEvent":  true,
		"IssuesEvent":       true,
		"CreateEvent":       true,
		"DeleteEvent":       true,
		"ForkEvent":         true,
		"WatchEvent":        true,
		"ReleaseEvent":      true,
		"MemberEvent":       true,
		"PublicEvent":       true,
		"GollumEvent":       true,
		"IssueCommentEvent": true,
		"LabelEvent":        true,
		"MilestoneEvent":    true,
	}
	return relevantEvents[typeOfEvent]
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
		if !checkEvents(event.Type) {
			continue
		}

		switch event.Type {
		case "PushEvent":
			fmt.Printf("📤 Pushed to %s\n", event.Repo.Name)
		case "PullRequestEvent":
			fmt.Printf("🔄 Opened a Pull Request in %s\n", event.Repo.Name)
		case "IssuesEvent":
			fmt.Printf("📝 Created or commented on an issue in %s\n", event.Repo.Name)
		case "CreateEvent":
			fmt.Printf("✨ Created a new repository: %s\n", event.Repo.Name)
		case "DeleteEvent":
			fmt.Printf("🗑️Deleted something in: %s\n", event.Repo.Name)
		case "ForkEvent":
			fmt.Printf("🍴 Forked the repository: %s\n", event.Repo.Name)
		case "WatchEvent":
			fmt.Printf("⭐ Started watching the repository: %s\n", event.Repo.Name)
		case "ReleaseEvent":
			fmt.Printf("🚀 Released a new version in: %s\n", event.Repo.Name)
		case "MemberEvent":
			fmt.Printf("👥 A member was added or removed in: %s\n", event.Repo.Name)
		case "PublicEvent":
			fmt.Printf("🌍 Made the repository public: %s\n", event.Repo.Name)
		case "GollumEvent":
			fmt.Printf("📖 Updated the wiki in: %s\n", event.Repo.Name)
		case "IssueCommentEvent":
			fmt.Printf("💬 Commented on an issue or pull request in: %s\n", event.Repo.Name)
		case "LabelEvent":
			fmt.Printf("🏷️ A label was added or modified in: %s\n", event.Repo.Name)
		case "MilestoneEvent":
			fmt.Printf("🎯 A milestone was created or updated in: %s\n", event.Repo.Name)
		}
	}
}

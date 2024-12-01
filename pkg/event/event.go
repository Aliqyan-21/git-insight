package event

import "fmt"

type PushPayload struct {
	Commits []struct {
		Message string `json:"message"`
	} `json:"commits"`
}

type Event struct {
	Type string `json:"type"`
	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`
	Payload PushPayload `json:"payload"`
}

func CheckEvents(typeOfEvent string) bool {
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

func PrintEvent(event Event) {
	switch event.Type {
	case "PushEvent":
		fmt.Printf("📤 Pushed to %s\n", event.Repo.Name)
		for _, commit := range event.Payload.Commits {
			fmt.Printf(" ~ %s\n", commit.Message)
		}
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

package main

import (
	"fmt"
	"github.com/aliqyan-21/git-insight/internal/github"
	"github.com/aliqyan-21/git-insight/pkg/event"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: git-insight <username>")
		return
	}

	username := os.Args[1]
	events, err := github.FetchUserEvents(username)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Events fetched successfully: \n\n")
	for _, e := range events {
		if !event.CheckEvents(e.Type) {
			continue
		}
		event.PrintEvent(e)
	}
}

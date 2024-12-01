package github

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aliqyan-21/git-insight/pkg/event"
	"io"
	"net/http"
)

// fetchUserEvents function fetches the github events for the username provided to us
func FetchUserEvents(username string) ([]event.Event, error) {
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

	var events []event.Event
	err = json.Unmarshal(body, &events)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error parsing json: %s", err))
	}

	return events, nil
}

package client

import (
	"encoding/json"
	"fmt"
)

// Notification is ...
type Notification struct {
	space  string
	apiKey string
}

// NewNotification is ...
func NewNotification(space string, apiKey string) *Notification {
	n := new(Notification)
	n.space = space
	n.apiKey = apiKey

	return n
}

// List function returns space notification.
func (n *Notification) List() ([]NotificationResponse, error) {
	api := "api/v2/notifications"
	cli := NewClient(n.space, n.apiKey)
	body, err := cli.Get(api, nil)
	if err != nil {
		return nil, err
	}

	var r []NotificationResponse
	if err = json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	return r, nil
}

// PrintCSV function prints list of notification in CSV format.
func (n *Notification) PrintCSV(r []NotificationResponse) {
	for _, n := range r {
		fmt.Printf("%d,%s,%s\n",
			n.ID,
			n.Issue.IssueKey,
			n.Issue.Summary)
	}
}

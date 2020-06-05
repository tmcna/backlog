package client

import (
	"encoding/json"
	"fmt"
)

// Priority is ...
type Priority struct {
	space  string
	apiKey string
}

// NewPriority is ...
func NewPriority(space string, apiKey string) *Priority {
	c := new(Priority)
	c.space = space
	c.apiKey = apiKey

	return c
}

// List function returns list of priorities.
func (c *Priority) List() ([]PriorityResponse, error) {
	api := "api/v2/priorities"
	cli := NewClient(c.space, c.apiKey)
	body, err := cli.Get(api, nil)
	if err != nil {
		return nil, err
	}

	var r []PriorityResponse
	if err = json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	return r, nil
}

// GetID function gets id from Priority name.
func (c *Priority) GetID(name string) (int, error) {
	api := "api/v2/priorities"
	cli := NewClient(c.space, c.apiKey)
	body, err := cli.Get(api, nil)
	if err != nil {
		return -1, err
	}

	var r []PriorityResponse
	if err = json.Unmarshal(body, &r); err != nil {
		return -1, err
	}

	for _, n := range r {
		if n.Name == name {
			return n.ID, nil
		}
	}

	return -1, nil
}

// PrintCSV function prints list of priorities in CSV format.
func (c *Priority) PrintCSV(r []PriorityResponse) {
	for _, n := range r {
		fmt.Printf("%d,%s\n", n.ID, n.Name)
	}
}

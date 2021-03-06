package client

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// Status is ...
type Status struct {
	space      string
	apiKey     string
	projectKey string
	request    url.Values
}

// NewStatus is constructor.
func NewStatus(space string, apiKey string, projectKey string) *Status {
	t := new(Status)
	t.space = space
	t.apiKey = apiKey
	t.projectKey = projectKey
	t.request = url.Values{}

	return t
}

// Request is ...
func (t *Status) Request(key string, value string) {
	t.request.Add(key, value)
}

// List function returns list of statuses.
func (t *Status) List() ([]StatusResponse, error) {
	api := fmt.Sprintf("api/v2/projects/%s/statuses", t.projectKey)

	cli := NewClient(t.space, t.apiKey)
	body, err := cli.Get(api, nil)
	if err != nil {
		return nil, err
	}

	var r []StatusResponse
	if err = json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	return r, nil
}

// Add function adds new Status to the project.
func (t *Status) Add() (*StatusResponse, error) {
	api := fmt.Sprintf("api/v2/projects/%s/statuses", t.projectKey)

	cli := NewClient(t.space, t.apiKey)
	body, err := cli.Post(api, t.request)
	if err != nil {
		return nil, err
	}
	var r StatusResponse
	if err = json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

// Delete function deletes status.
func (t *Status) Delete(statusID int, substituteStatusID int) (*StatusResponse, error) {
	api := fmt.Sprintf("api/v2/projects/%s/statuses/%s", t.projectKey, strconv.Itoa(statusID))

	values := url.Values{}
	values.Set("substituteStatusId", strconv.Itoa(substituteStatusID))
	cli := NewClient(t.space, t.apiKey)
	body, err := cli.Delete(api, values)
	if err != nil {
		return nil, err
	}

	var r StatusResponse
	if err = json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

// GetID function gets id from Status name.
func (t *Status) GetID(name string) (int, error) {
	api := fmt.Sprintf("api/v2/projects/%s/statuses", t.projectKey)

	cli := NewClient(t.space, t.apiKey)
	body, err := cli.Get(api, nil)
	if err != nil {
		return -1, err
	}

	var r []StatusResponse
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

// PrintCSV function prints list of statuses in CSV format.
func (t *Status) PrintCSV(r []StatusResponse) {
	for _, n := range r {
		fmt.Printf("%d,%s\n", n.ID, n.Name)
	}
}

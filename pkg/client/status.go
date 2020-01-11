package client

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// Status is ...
type Status struct {
	space   string
	apiKey  string
	request url.Values
}

// NewStatus is constructor.
func NewStatus(space string, apiKey string) *Status {
	t := new(Status)
	t.space = space
	t.apiKey = apiKey
	t.request = url.Values{}

	return t
}

// Request is ...
func (t *Status) Request(key string, value string) {
	t.request.Add(key, value)
}

// List function returns list of statuses.
func (t *Status) List(projectKey string) ([]StatusResponse, error) {
	api := "api/v2/projects/" + projectKey + "/statuses"
	cli := NewClient(t.space, t.apiKey)
	body, err := cli.Get(api, nil)
	if err != nil {
		return nil, err
	}

	var r []StatusResponse
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Add function adds new Status to the project.
func (t *Status) Add(projectKey string) (*StatusResponse, error) {

	api := "api/v2/projects/" + projectKey + "/statuses"

	cli := NewClient(t.space, t.apiKey)
	body, err := cli.Post(api, t.request)
	if err != nil {
		return nil, err
	}
	var r StatusResponse
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// Delete function deletes status.
func (t *Status) Delete(projectKey string, statusID int, substituteStatusID int) (*StatusResponse, error) {

	api := "api/v2/projects/" + projectKey + "/statuses/" + strconv.Itoa(statusID)
	values := url.Values{}
	values.Set("substituteStatusId", strconv.Itoa(substituteStatusID))
	cli := NewClient(t.space, t.apiKey)
	body, err := cli.Delete(api, values)
	if err != nil {
		return nil, err
	}

	var r StatusResponse
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// GetID function gets id from Status name.
func (t *Status) GetID(projectKey string, name string) (int, error) {
	api := "api/v2/projects/" + projectKey + "/statuses"
	cli := NewClient(t.space, t.apiKey)
	body, err := cli.Get(api, nil)
	if err != nil {
		return -1, err
	}

	var r []StatusResponse
	err = json.Unmarshal(body, &r)
	if err != nil {
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
	for i := 0; i < len(r); i++ {
		fmt.Printf("%d,%s\n", r[i].ID, r[i].Name)
	}
}

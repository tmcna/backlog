package client

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// IssueType is ...
type IssueType struct {
	space      string
	apiKey     string
	projectKey string
	request    url.Values
}

// NewIssueType is constructor.
func NewIssueType(space string, apiKey string, projectKey string) *IssueType {
	t := new(IssueType)
	t.space = space
	t.apiKey = apiKey
	t.projectKey = projectKey
	t.request = url.Values{}

	return t
}

// Request is ...
func (t *IssueType) Request(key string, value string) {
	t.request.Add(key, value)
}

// List function returns list of Issue Types in the project.
func (t *IssueType) List() ([]IssueTypeResponse, error) {
	api := fmt.Sprintf("api/v2/projects/%s/issueTypes", t.projectKey)

	cli := NewClient(t.space, t.apiKey)
	body, err := cli.Get(api, nil)
	if err != nil {
		return nil, err
	}

	var r []IssueTypeResponse
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Add function adds new Issue Type to the project.
func (t *IssueType) Add() (*IssueTypeResponse, error) {
	api := fmt.Sprintf("api/v2/projects/%s/issueTypes", t.projectKey)

	cli := NewClient(t.space, t.apiKey)
	body, err := cli.Post(api, t.request)
	if err != nil {
		return nil, err
	}
	var r IssueTypeResponse
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// Delete function deletes Issue Type.
func (t *IssueType) Delete(issueTypeID int, substituteIssueTypeID int) (*IssueTypeResponse, error) {
	api := fmt.Sprintf("api/v2/projects/%s/issueTypes/%s", t.projectKey, strconv.Itoa(issueTypeID))

	values := url.Values{}
	values.Set("substituteIssueTypeId", strconv.Itoa(substituteIssueTypeID))
	cli := NewClient(t.space, t.apiKey)
	body, err := cli.Delete(api, values)
	if err != nil {
		return nil, err
	}

	var r IssueTypeResponse
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// GetID function gets id from Issue Type name.
func (t *IssueType) GetID(name string) (int, error) {
	api := fmt.Sprintf("api/v2/projects/%s/issueTypes", t.projectKey)
	
	cli := NewClient(t.space, t.apiKey)
	body, err := cli.Get(api, nil)
	if err != nil {
		return -1, err
	}

	var r []IssueTypeResponse
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

// PrintCSV function prints list of Issue Type in CSV format.
func (t *IssueType) PrintCSV(r []IssueTypeResponse) {
	for i := 0; i < len(r); i++ {
		fmt.Printf("%d,%s\n", r[i].ID, r[i].Name)
	}
}

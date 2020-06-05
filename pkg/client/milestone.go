package client

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// Milestone is ...
type Milestone struct {
	space      string
	apiKey     string
	projectKey string
	request    url.Values
}

// NewMilestone is ...
func NewMilestone(space string, apiKey string, projectKey string) *Milestone {
	m := new(Milestone)
	m.space = space
	m.apiKey = apiKey
	m.projectKey = projectKey
	m.request = url.Values{}

	return m
}

// Request is ...
func (m *Milestone) Request(key string, value string) {
	m.request.Set(key, value)
}

// List function returns list of Versions/Milestones in the project.
func (m *Milestone) List() ([]MilestoneResponse, error) {
	api := fmt.Sprintf("api/v2/projects/%s/versions", m.projectKey)
	
	cli := NewClient(m.space, m.apiKey)
	body, err := cli.Get(api, nil)
	if err != nil {
		return nil, err
	}

	var r []MilestoneResponse
	if err = json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	return r, nil
}

// Add function adds new Version/Milestone to the project.
func (m *Milestone) Add() (*MilestoneResponse, error) {
	api := fmt.Sprintf("api/v2/projects/%s/versions", m.projectKey)

	cli := NewClient(m.space, m.apiKey)
	body, err := cli.Post(api, m.request)
	if err != nil {
		return nil, err
	}

	var r MilestoneResponse
	if err = json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

// Delete function deletes Milestone.
func (m *Milestone) Delete(id int) (*MilestoneResponse, error) {
	api := fmt.Sprintf("api/v2/projects/%s/versions/%s", m.projectKey, strconv.Itoa(id))

	cli := NewClient(m.space, m.apiKey)
	body, err := cli.Delete(api, nil)
	if err != nil {
		return nil, err
	}

	var r MilestoneResponse
	if err = json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

// GetID function gets id from Version/Milestone name.
func (m *Milestone) GetID(projectKey string, name string) (int, error) {
	api := fmt.Sprintf("api/v2/projects/%s/versions", projectKey)

	cli := NewClient(m.space, m.apiKey)
	body, err := cli.Get(api, nil)
	if err != nil {
		return -1, err
	}

	var r []MilestoneResponse
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

// PrintCSV function prints list of Milestones in CSV format.
func (m *Milestone) PrintCSV(r []MilestoneResponse) {
	for _, n := range r {
		fmt.Printf("%d,%d,%s,%s,%s,%s,%d\n",
			n.ID,
			n.ProjectID,
			n.Name,
			n.Description,
			n.StartDate,
			n.ReleaseDueDate,
			n.DisplayOrder)
	}
}

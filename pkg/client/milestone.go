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
	api := "api/v2/projects/" + m.projectKey + "/versions"
	cli := NewClient(m.space, m.apiKey)
	body, err := cli.Get(api, nil)
	if err != nil {
		return nil, err
	}

	var r []MilestoneResponse
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// Add function adds new Version/Milestone to the project.
func (m *Milestone) Add() (*MilestoneResponse, error) {

	api := "api/v2/projects/" + m.projectKey + "/versions"
	cli := NewClient(m.space, m.apiKey)
	body, err := cli.Post(api, m.request)
	if err != nil {
		return nil, err
	}

	var r MilestoneResponse
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// Delete function deletes Milestone.
func (m *Milestone) Delete(id int) (*MilestoneResponse, error) {

	api := "api/v2/projects/" + m.projectKey + "/versions/" + strconv.Itoa(id)

	cli := NewClient(m.space, m.apiKey)
	body, err := cli.Delete(api, nil)
	if err != nil {
		return nil, err
	}

	var r MilestoneResponse
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// GetID function gets id from Version/Milestone name.
func (m *Milestone) GetID(projectKey string, name string) (int, error) {
	api := "api/v2/projects/" + projectKey + "/versions"
	cli := NewClient(m.space, m.apiKey)
	body, err := cli.Get(api, nil)
	if err != nil {
		return -1, err
	}

	var r []MilestoneResponse
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

// PrintCSV function prints list of Milestones in CSV format.
func (m *Milestone) PrintCSV(r []MilestoneResponse) {
	for i := 0; i < len(r); i++ {
		fmt.Printf("%d,%d,%s,%s,%s,%s,%d\n",
			r[i].ID,
			r[i].ProjectID,
			r[i].Name,
			r[i].Description,
			r[i].StartDate,
			r[i].ReleaseDueDate,
			r[i].DisplayOrder)
	}
}

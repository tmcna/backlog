package client

import (
	"encoding/json"
	"fmt"
)

// Projects is ...
type Projects struct {
	space  string
	apiKey string
}

// Project is ...
type Project struct {
	space      string
	apiKey     string
	projectKey string
	project    *ProjectResponse
	priority   []PriorityResponse
	issueType  []IssueTypeResponse
	user       []UserResponse
	category   []CategoryResponse
	milestone  []MilestoneResponse
	status     []StatusResponse
}

// NewProjects is ...
func NewProjects(space string, apiKey string) *Projects {
	p := new(Projects)
	p.space = space
	p.apiKey = apiKey

	return p
}

// List function returns list of projects.
func (p *Projects) List() ([]ProjectResponse, error) {
	api := "api/v2/projects"

	cli := NewClient(p.space, p.apiKey)
	body, err := cli.Get(api, nil)
	if err != nil {
		return nil, err
	}

	var r []ProjectResponse
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// PrintCSV function prints list of projects in CSV format.
func (p *Projects) PrintCSV(r []ProjectResponse) {
	for _, n := range r {
		fmt.Printf("%d,%s,%s\n", n.ID, n.ProjectKey, n.Name)
	}
}

// NewProject is ...
func NewProject(space string, apiKey string, projectKey string) (*Project, error) {
	p := new(Project)
	p.space = space
	p.apiKey = apiKey
	p.projectKey = projectKey
	err := p.Reset()
	if err != nil {
		return nil, err
	}

	return p, nil
}

// Info function returns information about project.
func (p *Project) Info() (*ProjectResponse, error) {
	api := "api/v2/projects/" + p.projectKey

	cli := NewClient(p.space, p.apiKey)
	body, err := cli.Get(api, nil)
	if err != nil {
		return nil, err
	}

	var r ProjectResponse
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// Print function prints project informations.
func (p *Project) Print() {
	fmt.Printf("%d,%s,%s\n", p.project.ID, p.projectKey, p.project.Name)
	for _, n := range p.issueType {
		fmt.Printf("IssueType: %d,%s,%s,%d\n", n.ID, n.Name, n.Color, n.DisplayOrder)
	}
	for _, n := range p.user {
		fmt.Printf("User: %d,%s,%s,%d,%s,%s\n", n.ID, n.UserID, n.Name, n.RoleType, n.Lang, n.MailAddress)
	}
	for _, n := range p.category {
		fmt.Printf("Category: %d,%s,%d\n", n.ID, n.Name, n.DisplayOrder)
	}
	for _, n := range p.milestone {
		fmt.Printf("Milestone: %d,%s,%s,%s,%s,%d\n", n.ID, n.Name, n.Description, n.StartDate, n.ReleaseDueDate, n.DisplayOrder)
	}
	for _, n := range p.status {
		fmt.Printf("Status: %d,%s,%s,%d\n", n.ID, n.Name, n.Color, n.DisplayOrder)
	}
}

// GetID function gets id from project.
func (p *Project) GetID() int {
	return p.project.ID
}

// Reset function resets project infomation.
func (p *Project) Reset() error {
	var err error
	// get project informations.
	p.project, err = p.Info()
	if err != nil {
		return err
	}

	// get issue type in the project.
	issueType := NewIssueType(p.space, p.apiKey, p.projectKey)
	p.issueType, err = issueType.List()
	if err != nil {
		return err
	}

	// get list of project members.
	u := NewUser(p.space, p.apiKey)
	p.user, err = u.ListOfProject(p.projectKey)
	if err != nil {
		return err
	}

	// get priorities in the project.
	py := NewPriority(p.space, p.apiKey)
	p.priority, err = py.List()
	if err != nil {
		return err
	}

	// get categories in the project.
	c := NewCategory(p.space, p.apiKey, p.projectKey)
	p.category, err = c.List()
	if err != nil {
		return err
	}

	// get milestones in the project.
	milestone := NewMilestone(p.space, p.apiKey, p.projectKey)
	p.milestone, err = milestone.List()
	if err != nil {
		return err
	}

	// get statuses in the project.
	status := NewStatus(p.space, p.apiKey, p.projectKey)
	p.status, err = status.List()
	if err != nil {
		return err
	}

	return nil
}

// GetIssueTypeID function gets from Issue Type name.
func (p *Project) GetIssueTypeID(name string) int {
	var id int = -1
	for _, n := range p.issueType {
		if n.Name == name {
			id = n.ID
		}
	}
	return id
}

// GetPriorityID function gets from Priority name.
func (p *Project) GetPriorityID(name string) int {
	var id int = -1
	for _, n := range p.priority {
		if n.Name == name {
			id = n.ID
		}
	}
	return id
}

// GetUserID function gets from Priority name.
func (p *Project) GetUserID(name string) int {
	var id int = -1
	for _, n := range p.user {
		if n.Name == name {
			id = n.ID
		}
	}
	return id
}

// GetStatusID function gets from Priority name.
func (p *Project) GetStatusID(name string) int {
	var id int = -1
	for _, n := range p.status {
		if n.Name == name {
			id = n.ID
		}
	}
	return id
}

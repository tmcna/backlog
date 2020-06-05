package client

import (
	"encoding/json"
	"fmt"
)

// Space is ...
type Space struct {
	space  string
	apiKey string
}

// NewSpace is constructor.
func NewSpace(space string, apiKey string) *Space {
	p := new(Space)
	p.space = space
	p.apiKey = apiKey

	return p
}

// GetSpaceUsage function returns space usage.
func (p *Space) GetSpaceUsage() (*SpaceUsageResponse, error) {
	api := "api/v2/space/diskUsage"

	cli := NewClient(p.space, p.apiKey)
	body, err := cli.Get(api, nil)
	if err != nil {
		return nil, err
	}

	var r SpaceUsageResponse
	if err = json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

// GetProject function returns the Project object.
func (p *Space) GetProject(projectKey string) (*Project, error) {

	// プロジェクトの生成
	project, err := NewProject(p.space, p.apiKey, projectKey)
	if err != nil {
		return nil, err
	}

	return project, nil
}

// GetIssues function returns the Issue object.
func (p *Space) GetIssues() *Issue {

	i := NewIssue(p.space, p.apiKey)
	return i
}

// GetActivity function returns the Activity object.
func (p *Space) GetActivity(dispnum int, order int) *Activity {

	return NewActivity(p.space, p.apiKey, dispnum, order)
}

// PrintSpaceUsageCSV function prints space usage in CSV format.
func (p *Space) PrintSpaceUsageCSV(r *SpaceUsageResponse) {
	fmt.Printf("%s,%s,%s,%s,%s,%s,%s\n",
		"Capacity", "Issue", "Wiki", "File", "Subversion", "Git", "GitLFS")
	fmt.Printf("%d,%d,%d,%d,%d,%d,%d\n",
		r.Capacity, r.Issue, r.Wiki, r.File, r.Subversion, r.Git, r.GitLFS)
}

package backlogcli

import (
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/BurntSushi/toml"
	"github.com/tealeg/xlsx"

	"github.com/tmcna/backlog/pkg/client"
)

// Subcommand struct
type Subcommand struct {
	space  string
	apiKey string
}

// NewSubcommand is a constructor.
func NewSubcommand() *Subcommand {
	s := new(Subcommand)
	cfg, err := client.NewConfig()
	if err != nil {
		return nil
	}
	s.space, s.apiKey = cfg.Setup()
	return s
}

// UserList function user list subcommand.
func (s *Subcommand) UserList() error {
	user := client.NewUser(s.space, s.apiKey)
	r, err := user.List()
	if err != nil {
		return err
	}
	user.PrintCSV(r)
	return nil
}

// IssueList function issue list subcommand.
func (s *Subcommand) IssueList() error {
	issue := client.NewIssue(s.space, s.apiKey)
	for issue.HasNext() {
		r, err := issue.List()
		if err != nil {
			return err
		}
		issue.PrintCSV(r)
	}
	return nil
}

// IssueAddRequest structure　for TOML file.
type IssueAddRequest struct {
	ProjectKey     string `toml:"projectKey"`
	Summary        string `toml:"summary"`
	Description    string `toml:"description"`
	Assignee       string `toml:"asignee"`
	StartDate      string `toml:"startDate"`
	DueDate        string `toml:"dueDate"`
	EstimatedHours string `toml:"estimatedHours"`
	IssueType      string `toml:"issueType"`
	Priority       string `toml:"priority"`
}

// IssueAdd function issue add subcommand.
func (s *Subcommand) IssueAdd(path string) error {

	e := filepath.Ext(path)

	var q IssueAddRequest

	switch e {
	case ".toml":
		// Toml
		_, err := toml.DecodeFile(path, &q)
		if err != nil {
			return err
		}
	case ".xlsx":
		// Excel
		excel, err := xlsx.OpenFile(path)
		if err != nil {
			return err
		}

		sheet := excel.Sheets[0]
		q.ProjectKey = sheet.Rows[1].Cells[0].Value     //ProjectKey
		q.Summary = sheet.Rows[1].Cells[1].Value        //Summary
		q.Description = sheet.Rows[1].Cells[2].Value    //Description
		q.Assignee = sheet.Rows[1].Cells[3].Value       //Assignee
		q.StartDate = sheet.Rows[1].Cells[4].Value      //StartDate
		q.DueDate = sheet.Rows[1].Cells[5].Value        //DueDate
		q.EstimatedHours = sheet.Rows[1].Cells[6].Value //EstimateHours
		q.IssueType = sheet.Rows[1].Cells[7].Value      //IssueTypeID
		q.Priority = sheet.Rows[1].Cells[8].Value       //PriorityID
	default:
		err := fmt.Errorf("file type error.file=%s", path)
		return err
	}

	// check required parameter.
	if q.ProjectKey == "" {
		return fmt.Errorf("Error:none required parameter %s", "projectKey")
	}
	if q.Summary == "" {
		return fmt.Errorf("Error:none required parameter %s", "summary")
	}
	if q.IssueType == "" {
		return fmt.Errorf("Error:none required parameter %s", "issueType")
	}
	if q.Priority == "" {
		return fmt.Errorf("Error:none required parameter %s", "proprity")
	}

	cfg, err := client.NewConfig()
	if err != nil {
		return err
	}
	space, apiKey := cfg.Setup()

	sp := client.NewSpace(space, apiKey)
	p, err := sp.GetProject(q.ProjectKey)
	if err != nil {
		return err
	}

	// Create request parameters for Backlog API.
	request := client.NewIssueRequest()
	request.ProjectID(strconv.Itoa(p.GetID()))                       //required parameter
	request.Summary(q.Summary)                                       //required parameter
	request.IssueTypeID(strconv.Itoa(p.GetIssueTypeID(q.IssueType))) //required parameter
	request.PriorityID(strconv.Itoa(p.GetPriorityID(q.Priority)))    //required parameter
	request.Description(q.Description)
	if q.Assignee != "" {
		request.Assignee(strconv.Itoa(p.GetUserID(q.Assignee)))
	}
	if q.StartDate != "" {
		request.StartDate(q.StartDate)
	}
	if q.DueDate != "" {
		request.DueDate(q.DueDate)
	}
	if q.EstimatedHours != "" {
		request.EstimatedHours(q.EstimatedHours)
	}
	issue := client.NewIssue(space, apiKey)
	r, err := issue.Add(request)
	if err != nil {
		return err
	}

	fmt.Printf("Add issueKey:%s\n", r.IssueKey)

	return nil
}

// IssueUpdateRequest structure　for TOML file.
type IssueUpdateRequest struct {
	ProjectKey     string `toml:"projectKey"`
	Summary        string `toml:"summary"`
	Description    string `toml:"description"`
	Status         string `toml:"status"`
	Assignee       string `toml:"asignee"`
	StartDate      string `toml:"startDate"`
	DueDate        string `toml:"dueDate"`
	EstimatedHours string `toml:"estimatedHours"`
	IssueType      string `toml:"issueType"`
	Priority       string `toml:"priority"`
	Comment        string `toml:"comment"`
}

// IssueUpdate function issue update subcommand.
func (s *Subcommand) IssueUpdate(optFile string, optStatus string, optAssignee string, optComment string, issueKey string) error {

	var q IssueUpdateRequest

	if optFile != "" {
		e := filepath.Ext(optFile)

		switch e {
		case ".toml":
			// Toml
			_, err := toml.DecodeFile(optFile, &q)
			if err != nil {
				return err
			}
		case ".xlsx":
			// Excel
			excel, err := xlsx.OpenFile(optFile)
			if err != nil {
				return err
			}

			sheet := excel.Sheets[0]
			q.ProjectKey = sheet.Rows[1].Cells[0].Value     //ProjectKey
			q.Summary = sheet.Rows[1].Cells[1].Value        //Summary
			q.Description = sheet.Rows[1].Cells[2].Value    //Description
			q.Assignee = sheet.Rows[1].Cells[3].Value       //Assignee
			q.StartDate = sheet.Rows[1].Cells[4].Value      //StartDate
			q.DueDate = sheet.Rows[1].Cells[5].Value        //DueDate
			q.EstimatedHours = sheet.Rows[1].Cells[6].Value //EstimateHours
			q.IssueType = sheet.Rows[1].Cells[7].Value      //IssueTypeID
			q.Priority = sheet.Rows[1].Cells[8].Value       //PriorityID
			q.Status = sheet.Rows[1].Cells[9].Value         //Status
		default:
			err := fmt.Errorf("file type error.file=%s", optFile)
			return err
		}
	}

	cfg, err := client.NewConfig()
	if err != nil {
		return err
	}
	space, apiKey := cfg.Setup()

	i2 := client.NewIssue(space, apiKey)
	r2, err := i2.Info(issueKey)
	if err != nil {
		return err
	}
	projectID := r2.ProjectID

	ps := client.NewProjects(space, apiKey)
	pr, err := ps.List()
	if err != nil {
		return err
	}
	var projectKey string
	for _, n := range pr {
		if n.ID == projectID {
			projectKey = n.ProjectKey
		}
	}

	sp := client.NewSpace(space, apiKey)
	p, err := sp.GetProject(projectKey)
	if err != nil {
		return err
	}

	// Create request parameters for Backlog API.
	request := client.NewIssueRequest()
	if q.Summary != "" {
		request.Summary(q.Summary)
	}
	if q.IssueType != "" {
		request.IssueTypeID(strconv.Itoa(p.GetIssueTypeID(q.IssueType)))
	}
	if q.Priority != "" {
		request.PriorityID(strconv.Itoa(p.GetPriorityID(q.Priority)))
	}
	request.Description(q.Description)
	if q.Assignee != "" {
		request.Assignee(strconv.Itoa(p.GetUserID(q.Assignee)))
	}
	if q.Status != "" {
		request.StatusID(strconv.Itoa(p.GetStatusID(q.Status)))
	}
	if q.StartDate != "" {
		request.StartDate(q.StartDate)
	}
	if q.DueDate != "" {
		request.DueDate(q.DueDate)
	}
	if q.EstimatedHours != "" {
		request.EstimatedHours(q.EstimatedHours)
	}
	if q.Comment != "" {
		request.Comment(q.Comment)
	}

	// Option
	if optStatus != "" {
		request.StatusID(strconv.Itoa(p.GetStatusID(optStatus)))
	}

	if optAssignee != "" {
		request.Assignee(strconv.Itoa(p.GetUserID(optAssignee)))
	}

	if optComment != "" {
		request.Comment(optComment)
	}

	// Issueオブジェクトを作成し、リクエストパラメーターを設定する。
	issue := client.NewIssue(space, apiKey)
	r, err := issue.Update(request, issueKey)
	if err != nil {
		return err
	}

	fmt.Printf("Update issueKey:%s\n", r.IssueKey)

	return nil
}

// CommentAdd function comment add subcommand.
func (s *Subcommand) CommentAdd(issueKey string, content string) error {
	cfg, err := client.NewConfig()
	if err != nil {
		return err
	}
	space, apiKey := cfg.Setup()

	// Create request parameters for Backlog API.
	q := client.NewCommentRequest()
	q.Content(content)

	// Commentオブジェクトを作成し、リクエストパラメーターを設定する。
	comment := client.NewComment(space, apiKey, issueKey)
	_, err = comment.Add(q)
	if err != nil {
		return err
	}

	return nil
}

// IssueDelete function issue delete subcommand.
func (s *Subcommand) IssueDelete(issueKey string) error {
	if issueKey == "" {
		return fmt.Errorf("Error: %s", "Issue key is not found.")
	}
	issue := client.NewIssue(s.space, s.apiKey)
	_, err := issue.Delete(issueKey)
	if err != nil {
		return err
	}
	return nil
}

// ActivityList function executes act subcommand.
func (s *Subcommand) ActivityList() error {
	act := client.NewActivity(s.space, s.apiKey, 32, client.DisplayOrderDesc)
	for act.HasNext() {
		r, err := act.List()
		if err != nil {
			return err
		}
		act.PrintCSV(r)
	}
	return nil
}

// NotifyList function executes notify subcommand.
func (s *Subcommand) NotifyList() error {
	n := client.NewNotification(s.space, s.apiKey)
	r, err := n.List()
	if err != nil {
		return err
	}
	n.PrintCSV(r)
	return nil
}

// ProjectList function executes project list subcommand.
func (s *Subcommand) ProjectList() error {
	p := client.NewProjects(s.space, s.apiKey)
	r, err := p.List()
	if err != nil {
		return err
	}
	p.PrintCSV(r)
	return nil
}

// ProjectInfo function executes project info subcommand.
func (s *Subcommand) ProjectInfo(projectKey string) error {
	p, err := client.NewProject(s.space, s.apiKey, projectKey)
	if err != nil {
		return err
	}
	p.Print()
	return nil
}

// SpaceUsage function executes space subcommand.
func (s *Subcommand) SpaceUsage() error {
	sp := client.NewSpace(s.space, s.apiKey)
	r, err := sp.GetSpaceUsage()
	if err != nil {
		return err
	}
	sp.PrintSpaceUsageCSV(r)
	return nil
}

package client

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// IssueRequest is ...
type IssueRequest struct {
	values url.Values
}

// NewIssueRequest is ...
func NewIssueRequest() *IssueRequest {
	q := new(IssueRequest)
	q.values = url.Values{}

	return q
}

// ProjectID function sets request parameters in the Request structure.
func (q *IssueRequest) ProjectID(value string) {
	q.values.Set("projectId", value)
}

// Summary function sets request parameters in the Request structure.
func (q *IssueRequest) Summary(value string) {
	q.values.Set("summary", value)
}

// Description function sets request parameters in the Request structure.
func (q *IssueRequest) Description(value string) {
	q.values.Set("description", value)
}

// Assignee function sets request parameters in the Request structure.
func (q *IssueRequest) Assignee(value string) {
	q.values.Set("assigneeId", value)
}

// StatusID function sets request parameters in the Request structure.
func (q *IssueRequest) StatusID(value string) {
	q.values.Set("statusId", value)
}

// StartDate function sets request parameters in the Request structure.
func (q *IssueRequest) StartDate(value string) {
	q.values.Set("startDate", value)
}

// DueDate function sets request parameters in the Request structure.
func (q *IssueRequest) DueDate(value string) {
	q.values.Set("dueDate", value)
}

// EstimatedHours function sets request parameters in the Request structure.
func (q *IssueRequest) EstimatedHours(value string) {
	q.values.Set("estimatedHours", value)
}

// IssueTypeID function sets request parameters in the Request structure.
func (q *IssueRequest) IssueTypeID(value string) {
	q.values.Add("issueTypeId", value)
}

// PriorityID function sets request parameters in the Request structure.
func (q *IssueRequest) PriorityID(value string) {
	q.values.Set("priorityId", value)
}

// Comment function sets request parameters in the Request structure.
func (q *IssueRequest) Comment(value string) {
	q.values.Set("comment", value)
}

// Issue structure.
type Issue struct {
	count  int
	space  string
	apiKey string
	query  url.Values
	next   bool
	offset int
}

// NewIssue is ...
func NewIssue(space string, apiKey string) *Issue {
	issue := new(Issue)
	issue.space = space
	issue.apiKey = apiKey
	issue.query = url.Values{}
	issue.next = true
	issue.offset = 0

	return issue
}

// Query is ...
func (t *Issue) Query(key string, value string) {
	t.query.Add(key, value)
}

// List function returns list of issues.
func (t *Issue) List() ([]IssueResponse, error) {
	api := "api/v2/issues"

	count := IssueResponseCount                   //1回のAPIで取り出す件数
	t.query.Set("count", strconv.Itoa(count))     //課題取得数を追加(毎回100件)
	t.query.Set("offset", strconv.Itoa(t.offset)) //課題取得開始位置を指定
	cli := NewClient(t.space, t.apiKey)
	body, err := cli.Get(api, t.query)
	if err != nil {
		return nil, err
	}

	var r []IssueResponse
	if err := json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	// 取得できた課題数が、指定した数より小さければ全取得と判断し、nextフラグをfalseにする。
	// 取得数=指定数の場合一度空振りするが頻度が少ないため許容する。
	if len(r) < count {
		t.next = false
	}
	t.offset += len(r)

	return r, nil
}

// HasNext is ...
func (t *Issue) HasNext() bool {
	return t.next
}

// Info function returns information about issue.
func (t *Issue) Info(issueKey string) (*IssueResponse, error) {
	api := "api/v2/issues/" + issueKey

	cli := NewClient(t.space, t.apiKey)
	body, err := cli.Get(api, nil)
	if err != nil {
		return nil, err
	}

	var r IssueResponse
	if err := json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

// Add function adds new issue.
func (t *Issue) Add(q *IssueRequest) (*IssueResponse, error) {
	api := "api/v2/issues"

	cli := NewClient(t.space, t.apiKey)
	body, err := cli.Post(api, q.values)
	if err != nil {
		return nil, err
	}
	// defer
	var r IssueResponse
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// Update function adds new issue.
func (t *Issue) Update(q *IssueRequest, issueKey string) (*IssueResponse, error) {
	api := "api/v2/issues/" + issueKey

	cli := NewClient(t.space, t.apiKey)
	body, err := cli.Patch(api, q.values)
	if err != nil {
		return nil, err
	}
	// defer
	var r IssueResponse
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// Delete function deletes issue.
func (t *Issue) Delete(issueKey string) (*IssueResponse, error) {

	issues := NewIssue(t.space, t.apiKey)
	issue, err := issues.Info(issueKey)
	if err != nil {
		return nil, err
	}
	api := "api/v2/issues/" + strconv.Itoa(issue.ID)

	cli := NewClient(t.space, t.apiKey)
	body, err := cli.Delete(api, nil)
	if err != nil {
		return nil, err
	}

	var r IssueResponse
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// PrintCSV function prints list of issue in CSV format.
func (t *Issue) PrintCSV(r []IssueResponse) {
	for _, n := range r {
		fmt.Printf("%s,%s,%s,%s,%s,%s\n",
			n.IssueKey,
			n.IssueType.Name,
			n.Summary,
			n.Status.Name,
			n.Assignee.Name,
			n.DueDate)
	}
}

// PrintIssueCSV function prints information about issue in CSV format.
func (t *Issue) PrintIssueCSV(r *IssueResponse) {
	fmt.Printf("%s,%s,%s,%s,%s,%s\n",
		r.IssueKey,
		r.IssueType.Name,
		r.Summary,
		r.Status.Name,
		r.Assignee.Name,
		r.DueDate)
}

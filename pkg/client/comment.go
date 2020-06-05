package client

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

// CommentRequest - Request parameter for Comment API.
type CommentRequest struct {
	values url.Values
}

// NewCommentRequest is a constructor.
func NewCommentRequest() *CommentRequest {
	q := new(CommentRequest)
	q.values = url.Values{}

	return q
}

// Content sets request parameter value.
func (q *CommentRequest) Content(value string) {
	q.values.Set("content", value)
}

// Comment is a structure for Comment APIs.
type Comment struct {
	space    string
	apiKey   string
	issueKey string
}

// NewComment is constructor for Comment APIs.
func NewComment(space string, apiKey string, issueKey string) *Comment {
	comment := new(Comment)
	comment.space = space
	comment.apiKey = apiKey
	comment.issueKey = issueKey

	return comment
}

// Add function adds a comment to the issue.
func (t *Comment) Add(q *CommentRequest) (*CommentResponse, error) {
	api := fmt.Sprintf("api/v2/issues/%s/comments", t.issueKey)
	cli := NewClient(t.space, t.apiKey)
	body, err := cli.Post(api, q.values)
	if err != nil {
		return nil, err
	}

	var r CommentResponse
	if err = json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

// Update function updates content of comment.
func (t *Comment) Update(id string, q *CommentRequest) (*CommentResponse, error) {
	api := fmt.Sprintf("api/v2/issues/%s/comments/%s", t.issueKey, id)
	cli := NewClient(t.space, t.apiKey)
	body, err := cli.Patch(api, q.values)
	if err != nil {
		return nil, err
	}
	var r CommentResponse
	if err = json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

// List function returns list of comments in issue.
func (t *Comment) List() ([]CommentResponse, error) {
	api := fmt.Sprintf("api/v2/issues/%s/comments", t.issueKey)

	cli := NewClient(t.space, t.apiKey)
	body, err := cli.Get(api, nil)
	if err != nil {
		return nil, err
	}

	var r []CommentResponse
	if err = json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	return r, nil
}

// Delete function deletes Category by category ID.
func (t *Comment) Delete(id int) (*CommentResponse, error) {

	api := "api/v2/issues/" + t.issueKey + "/comments/" + strconv.Itoa(id)

	cli := NewClient(t.space, t.apiKey)
	body, err := cli.Delete(api, nil)
	if err != nil {
		return nil, err
	}

	var r CommentResponse
	if err = json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

// PrintCSV function prints list of comments in CSV format.
func (t *Comment) PrintCSV(r []CommentResponse) {
	for _, n := range r {
		fmt.Printf("%d,%s\n", n.ID, n.Content)
	}
}

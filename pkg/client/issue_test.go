package client

import (
	"fmt"
	"strconv"
	"testing"
)

func TestIssueListQuery(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}
	space, apiKey := cfg.Setup()

	issueType := NewIssueType(space, apiKey)
	taskID, err := issueType.GetID("TESTTOOL", "タスク")  //種別名から種別IDを求める
	otherID, err := issueType.GetID("TESTTOOL", "その他") //種別名から種別IDを求める

	issue := NewIssue(space, apiKey)
	issue.Query("issueTypeId[]", strconv.Itoa(taskID))
	issue.Query("issueTypeId[]", strconv.Itoa(otherID))
	r, err := issue.List()
	issue.PrintCSV(r)
}

func TestIssueList(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}
	space, apiKey := cfg.Setup()

	issue := NewIssue(space, apiKey)
	for issue.HasNext() {
		r, err := issue.List()
		if err != nil {
			t.Fatal(err)
		}
		issue.PrintCSV(r)
	}
}

func TestIssueInfo(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}
	space, apiKey := cfg.Setup()

	issue := NewIssue(space, apiKey)
	r, err := issue.Info("TESTTOOL-64")
	if err != nil {
		t.Fatal(err)
	}
	issue.PrintIssueCSV(r)
}

func TestIssueAddDelete(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}
	space, apiKey := cfg.Setup()

	p, err := NewProject(space, apiKey, "TESTTOOL")
	if err != nil {
		t.Fatal(err)
	}
	// Backlog APIのリクエストパラメーターを作成する。
	request := NewIssueRequest()
	request.ProjectID(strconv.Itoa(p.GetID()))
	request.Summary("課題削除のテスト")
	request.Description("課題の詳細")
	request.StartDate("2019-12-14")
	request.DueDate("2019-12-31")
	request.EstimatedHours("8")
	request.IssueTypeID("1626")
	request.PriorityID("1")

	// Issueオブジェクトを作成し、リクエストパラメーターを設定する。
	issue := NewIssue(space, apiKey)
	r, err := issue.Add(request)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Add issueKey:%s\n", r.IssueKey)

	_, err = issue.Delete(r.IssueKey)
	if err != nil {
		t.Fatal(err)
	}
}

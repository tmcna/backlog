package client

import (
	"fmt"
	"strconv"
	"testing"
)

func TestIssue_ListQuery(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}

	issueType := NewIssueType(cfg.Space, cfg.APIKey, "TESTTOOL")
	taskID, err := issueType.GetID("タスク")  //種別名から種別IDを求める
	otherID, err := issueType.GetID("その他") //種別名から種別IDを求める

	issue := NewIssue(cfg.Space, cfg.APIKey)
	issue.Query("issueTypeId[]", strconv.Itoa(taskID))
	issue.Query("issueTypeId[]", strconv.Itoa(otherID))
	r, err := issue.List()
	issue.PrintCSV(r)
}

func TestIssue_List(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}

	issue := NewIssue(cfg.Space, cfg.APIKey)
	for issue.HasNext() {
		r, err := issue.List()
		if err != nil {
			t.Fatal(err)
		}
		issue.PrintCSV(r)
	}
}

func TestIssue_Info(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}

	issue := NewIssue(cfg.Space, cfg.APIKey)
	r, err := issue.Info("TESTTOOL-64")
	if err != nil {
		t.Fatal(err)
	}
	issue.PrintIssueCSV(r)
}

func TestIssue_AddDelete(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}

	p, err := NewProject(cfg.Space, cfg.APIKey, "TESTTOOL")
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
	issue := NewIssue(cfg.Space, cfg.APIKey)
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

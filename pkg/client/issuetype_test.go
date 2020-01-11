package client

import (
	"fmt"
	"testing"
)

func TestIssueTypeList(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}
	space, apiKey := cfg.Setup()

	issueType := NewIssueType(space, apiKey)
	r, err := issueType.List("TESTTOOL")
	issueType.PrintCSV(r)
}

func TestIssueTypeAddDelete(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}
	space, apiKey := cfg.Setup()

	issueType := NewIssueType(space, apiKey)
	issueType.Request("name", "テスト")
	issueType.Request("color", "#e30000")
	result, err := issueType.Add("TESTTOOL")
	if err != nil {
		t.Fatal(err)
	}
	id := result.ID
	r, err := issueType.List("TESTTOOL")
	issueType.PrintCSV(r)
	_, err = issueType.Delete("TESTTOOL", id, 1626)
	if err != nil {
		t.Fatal(err)
	}
	r, err = issueType.List("TESTTOOL")
	issueType.PrintCSV(r)
}

func TestIssueTypeGetID(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}
	space, apiKey := cfg.Setup()

	issueType := NewIssueType(space, apiKey)
	id, err := issueType.GetID("TESTTOOL", "タスク")
	if err != nil {
		t.Fatal(err)
	}
	if id == -1 {
		t.Fatal(err)
	}
	fmt.Printf("IssueTypeID:%d", id)
}

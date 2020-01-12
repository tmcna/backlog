package client

import (
	"fmt"
	"testing"
)

func TestIssueType_List(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}

	issueType := NewIssueType(cfg.Space, cfg.APIKey, "TESTTOOL")
	r, err := issueType.List()
	issueType.PrintCSV(r)
}

func TestIssueType_AddDelete(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}

	issueType := NewIssueType(cfg.Space, cfg.APIKey, "TESTTOOL")
	issueType.Request("name", "テスト")
	issueType.Request("color", "#e30000")
	result, err := issueType.Add()
	if err != nil {
		t.Fatal(err)
	}
	id := result.ID
	r, err := issueType.List()
	issueType.PrintCSV(r)
	_, err = issueType.Delete(id, 1626)
	if err != nil {
		t.Fatal(err)
	}
	r, err = issueType.List()
	issueType.PrintCSV(r)
}

func TestIssueType_GetID(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}

	issueType := NewIssueType(cfg.Space, cfg.APIKey, "TESTTOOL")
	id, err := issueType.GetID("タスク")
	if err != nil {
		t.Fatal(err)
	}
	if id == -1 {
		t.Fatal(err)
	}
	fmt.Printf("IssueTypeID:%d", id)
}

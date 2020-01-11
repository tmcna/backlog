package client

import (
	"fmt"
	"testing"
)

func TestProjects(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}
	space, apiKey := cfg.Setup()

	p := NewProjects(space, apiKey)
	r, err := p.List()
	if err != nil {
		t.Fatal(err)
	}
	p.PrintCSV(r)
}

func TestProject(t *testing.T) {
	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}
	space, apiKey := cfg.Setup()

	p, err := NewProject(space, apiKey, "TESTTOOL")
	if err != nil {
		t.Fatal(err)
	}
	p.Print()
}

func TestProjectIssueTypeID(t *testing.T) {
	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}
	space, apiKey := cfg.Setup()

	p, err := NewProject(space, apiKey, "TESTTOOL")
	if err != nil {
		t.Fatal(err)
	}
	id := p.GetIssueTypeID("タスク")
	fmt.Println(id)
}

func TestProjectPriorityID(t *testing.T) {
	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}
	space, apiKey := cfg.Setup()

	p, err := NewProject(space, apiKey, "TESTTOOL")
	if err != nil {
		t.Fatal(err)
	}
	id := p.GetPriorityID("高")
	fmt.Println(id)
}

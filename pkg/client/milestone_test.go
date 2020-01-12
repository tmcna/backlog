package client

import (
	"fmt"
	"testing"
)

func TestMilestone_List(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}
	space, apiKey := cfg.Setup()

	m := NewMilestone(space, apiKey, "TESTTOOL")
	r, err := m.List()
	if err != nil {
		t.Fatal(err)
	}
	m.PrintCSV(r)
}

func TestMilestone_AddDelete(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}
	space, apiKey := cfg.Setup()

	var r *MilestoneResponse
	m := NewMilestone(space, apiKey, "TESTTOOL")
	m.Request("name", "テストバージョン")
	m.Request("description", "バージョンの説明")
	m.Request("startDate", "2019-01-01")
	m.Request("releaseDueDate", "2019-12-31")
	r, err = m.Add()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Milestone:%d\n", r.ID)

	r, err = m.Delete(r.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMilestone_GetID(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}
	space, apiKey := cfg.Setup()

	m := NewMilestone(space, apiKey, "TESTTOOL")
	id, err := m.GetID("TESTTOOL", "テストバージョン")
	if err != nil {
		t.Fatal(err)
	}
	if id == -1 {
		t.Fatal(err)
	}
	fmt.Printf("ID:%d\n", id)
}

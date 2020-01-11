package client

import (
	"testing"
)

func TestSpaceActivities(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}
	space, apiKey := cfg.Setup()

	s := NewSpace(space, apiKey)
	p, err := s.GetProject("TESTTOOL")
	if err != nil {
		t.Fatal(err)
	}
	p.Print()

	a := s.GetActivity(30, DisplayOrderAsc)
	for a.HasNext() {
		r, err := a.List()
		if err != nil {
			t.Fatal(err)
		}
		a.PrintCSV(r)
	}
}

func TestSpaceIssues(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}
	space, apiKey := cfg.Setup()

	s := NewSpace(space, apiKey)
	i := s.GetIssues()
	for i.HasNext() {
		r, err := i.List()
		if err != nil {
			t.Fatal(err)
		}
		i.PrintCSV(r)
	}
}

func TestSpaceUsage(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}
	space, apiKey := cfg.Setup()

	s := NewSpace(space, apiKey)
	r, err := s.GetSpaceUsage()
	if err != nil {
		t.Fatal(err)
	}
	s.PrintSpaceUsageCSV(r)
}

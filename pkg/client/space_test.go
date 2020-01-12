package client

import (
	"testing"
)

func TestSpace_GetActivities(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}

	s := NewSpace(cfg.Space, cfg.APIKey)
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

func TestSpace_GetIssues(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}

	s := NewSpace(cfg.Space, cfg.APIKey)
	i := s.GetIssues()
	for i.HasNext() {
		r, err := i.List()
		if err != nil {
			t.Fatal(err)
		}
		i.PrintCSV(r)
	}
}

func TestSpace_Usage(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}

	s := NewSpace(cfg.Space, cfg.APIKey)
	r, err := s.GetSpaceUsage()
	if err != nil {
		t.Fatal(err)
	}
	s.PrintSpaceUsageCSV(r)
}

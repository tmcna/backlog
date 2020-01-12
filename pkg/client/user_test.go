package client

import (
	"testing"
)

func TestUser_List(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}

	user := NewUser(cfg.Space, cfg.APIKey)
	r, err := user.List()
	user.PrintCSV(r)
}

func TestUser_ListOfProject(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}

	user := NewUser(cfg.Space, cfg.APIKey)
	r, err := user.ListOfProject("TESTTOOL")
	user.PrintCSV(r)
}

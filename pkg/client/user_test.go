package client

import (
	"testing"
)

func TestUser_List(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}
	space, apiKey := cfg.Setup()

	user := NewUser(space, apiKey)
	r, err := user.List()
	user.PrintCSV(r)
}

func TestUser_ListOfProject(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}
	space, apiKey := cfg.Setup()

	user := NewUser(space, apiKey)
	r, err := user.ListOfProject("TESTTOOL")
	user.PrintCSV(r)
}

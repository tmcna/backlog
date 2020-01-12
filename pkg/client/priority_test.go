package client

import (
	"fmt"
	"testing"
)

func TestPriority_List(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}

	c := NewPriority(cfg.Space, cfg.APIKey)
	r, err := c.List()
	c.PrintCSV(r)
}

func TestPriority_GetID(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}

	c := NewPriority(cfg.Space, cfg.APIKey)
	r, err := c.List()
	if err != nil {
		t.Fatal(err)
	}
	c.PrintCSV(r)
	id, err := c.GetID("高")
	if err != nil {
		t.Fatal(err)
	}
	if id == -1 {
		t.Fatal(err)
	}
	fmt.Printf("優先度ID:%d\n", id)
}

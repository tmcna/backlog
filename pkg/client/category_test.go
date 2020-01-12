package client

import (
	"fmt"
	"testing"
)

func TestCategory_List(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}
	space, apiKey := cfg.Setup()

	projectKey := "TESTTOOL"
	c := NewCategory(space, apiKey, projectKey)
	r, err := c.List()
	if err != nil {
		t.Fatal(err)
	}
	c.PrintCSV(r)
}

func TestCategory_AddDelete(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}
	space, apiKey := cfg.Setup()

	c := NewCategory(space, apiKey, "TESTTOOL")
	c.Request("name", "テストカテゴリー")
	r, err := c.Add()
	if err != nil {
		t.Fatal(err)
	}
	_, err = c.Delete(r.ID)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCategory_GetID(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}
	space, apiKey := cfg.Setup()

	c := NewCategory(space, apiKey, "TESTTOOL")
	r, err := c.List()
	if err != nil {
		t.Fatal(err)
	}
	c.PrintCSV(r)
	id, err := c.GetID("開発")
	if err != nil {
		t.Fatal(err)
	}
	if id == -1 {
		t.Fatal("Category ID not found.")
	}
	fmt.Printf("カテゴリーID:%d\n", id)
}

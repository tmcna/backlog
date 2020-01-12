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

	projectKey := "TESTTOOL"
	c := NewCategory(cfg.Space, cfg.APIKey, projectKey)
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

	c := NewCategory(cfg.Space, cfg.APIKey, "TESTTOOL")
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

	c := NewCategory(cfg.Space, cfg.APIKey, "TESTTOOL")
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

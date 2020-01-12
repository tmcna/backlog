package client

import (
	"fmt"
	"testing"
)

func TestStatus_List(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}

	status := NewStatus(cfg.Space, cfg.APIKey, "TESTTOOL")
	r, err := status.List()
	status.PrintCSV(r)
}

func TestStatus_Add(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}

	status := NewStatus(cfg.Space, cfg.APIKey, "TESTTOOL")
	status.Request("name", "回答待ち")
	status.Request("color", "#393939")
	_, err = status.Add()
	if err == nil { //フリープランでないときは err != nil にする
		t.Fatal(err)
	}
}

func TestStatus_GetID(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}

	status := NewStatus(cfg.Space, cfg.APIKey, "TESTTOOL")
	id, err := status.GetID("未対応")
	if err != nil {
		t.Fatal(err)
	}
	if id == -1 {
		t.Fatal(err)
	}
	fmt.Printf("StatusID:%d", id)
}

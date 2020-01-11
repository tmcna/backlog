package client

import (
	"fmt"
	"testing"
)

func TestStatusList(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}
	space, apiKey := cfg.Setup()

	status := NewStatus(space, apiKey)
	r, err := status.List("TESTTOOL")
	status.PrintCSV(r)
}

func TestStatusAdd(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}
	space, apiKey := cfg.Setup()

	status := NewStatus(space, apiKey)
	status.Request("name", "回答待ち")
	status.Request("color", "#393939")
	projectKey := "TESTTOOL"
	_, err = status.Add(projectKey)
	if err == nil { //フリープランでないときは err != nil にする
		t.Fatal(err)
	}
}

/* フリープランではテスト不可
func TestStatusDelete(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}
	space, apiKey := cfg.Setup()

	status := NewStatus(space, apiKey)
	status.Request("name", "回答待ち")
	status.Request("color", "#e30000")
	result, err := status.Add("TESTTOOL")
	if err != nil {
		t.Fatal(err)
	}
	id := result.ID
	var r []StatusResponse
	r, err = status.Find("TESTTOOL")
	status.Print(r)
	_, err = status.Delete("TESTTOOL", id, 1626)
	if err != nil {
		t.Fatal(err)
	}
	r, err = status.Find("TESTTOOL")
	status.Print(r)
}
*/

func TestStatusGetID(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}
	space, apiKey := cfg.Setup()

	status := NewStatus(space, apiKey)
	id, err := status.GetID("TESTTOOL", "未対応")
	if err != nil {
		t.Fatal(err)
	}
	if id == -1 {
		t.Fatal(err)
	}
	fmt.Printf("StatusID:%d", id)
}

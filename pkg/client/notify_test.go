package client

import (
	"testing"
)

func TestNotification_List(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}
	space, apiKey := cfg.Setup()

	n := NewNotification(space, apiKey)
	r, err := n.List()
	n.PrintCSV(r)
}

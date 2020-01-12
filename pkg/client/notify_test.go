package client

import (
	"testing"
)

func TestNotification_List(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}

	n := NewNotification(cfg.Space, cfg.APIKey)
	r, err := n.List()
	n.PrintCSV(r)
}

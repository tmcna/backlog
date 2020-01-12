package client

import (
	"testing"
)

func TestActivity_List(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}
	space, apiKey := cfg.Setup()

	act := NewActivity(space, apiKey, 50, DisplayOrderAsc)

	for act.HasNext() {
		r, err := act.List()
		if err != nil {
			t.Fatal(err)
		}
		act.PrintCSV(r)
	}

}

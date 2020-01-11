package client

import (
	"fmt"
	"testing"
)

func TestClient(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}
	space, apiKey := cfg.Setup()

	api := "api/v2" //Error

	cli := NewClient(space, apiKey)
	_, err = cli.Get(api, nil)
	if err == nil {
		t.Fatal(err)
	}
	s := fmt.Sprintf("%s", err)
	if s != "Error: StatusCode:400 Code:6 Message: Undefined resource. /api/v2/ MoreInfo:" {
		fmt.Println(err)
	}
}

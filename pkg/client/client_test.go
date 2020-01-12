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

	api := "api/v2" //Error

	cli := NewClient(cfg.Space, cfg.APIKey)
	_, err = cli.Get(api, nil)
	if err == nil {
		t.Fatal(err)
	}
	s := fmt.Sprintf("%s", err)
	if s != "Error: StatusCode:400 Code:6 Message: Undefined resource. /api/v2/ MoreInfo:" {
		fmt.Println(err)
	}
}

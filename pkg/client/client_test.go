package client

import (
	"fmt"
	"testing"
)

func TestClient_Get01(t *testing.T) {

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

func TestClient_Get02(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}

	api := "api/v2/issues"
	apiKey := "xxxxxxxxxxxxxxx" //Error

	cli := NewClient(cfg.Space, apiKey)
	_, err = cli.Get(api, nil)
	if err == nil {
		t.Fatal(err)
	}
	s := fmt.Sprintf("%s", err)
	if s != "Error: StatusCode:401 Code:11 Message: Authentication failure. MoreInfo:" {
		fmt.Println(err)
	}
}

func TestClient_Post01(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}

	api := "api/v2" //Error

	cli := NewClient(cfg.Space, cfg.APIKey)
	_, err = cli.Post(api, nil)
	if err == nil {
		t.Fatal(err)
	}
	s := fmt.Sprintf("%s", err)
	if s != "Error: StatusCode:400 Code:6 Message: Undefined resource. /api/v2/ MoreInfo:" {
		fmt.Println(err)
	}
}

func TestClient_Post02(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}

	api := "api/v2/issues"
	apiKey := "xxxxxxxxxxxxxxx" //Error

	cli := NewClient(cfg.Space, apiKey)
	_, err = cli.Post(api, nil)
	if err == nil {
		t.Fatal(err)
	}
	s := fmt.Sprintf("%s", err)
	if s != "Error: StatusCode:401 Code:11 Message: Authentication failure. MoreInfo:" {
		fmt.Println(err)
	}
}

func TestClient_Patch01(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}

	api := "api/v2" //Error

	cli := NewClient(cfg.Space, cfg.APIKey)
	_, err = cli.Patch(api, nil)
	if err == nil {
		t.Fatal(err)
	}
	s := fmt.Sprintf("%s", err)
	if s != "Error: StatusCode:400 Code:6 Message: Undefined resource. /api/v2/ MoreInfo:" {
		fmt.Println(err)
	}
}

func TestClient_Patch02(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}

	api := "api/v2/issues/0"
	apiKey := "xxxxxxxxxxxxxxx" //Error

	cli := NewClient(cfg.Space, apiKey)
	_, err = cli.Patch(api, nil)
	if err == nil {
		t.Fatal(err)
	}
	s := fmt.Sprintf("%s", err)
	if s != "Error: StatusCode:401 Code:11 Message: Authentication failure. MoreInfo:" {
		fmt.Println(err)
	}
}

func TestClient_Delete01(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}

	api := "api/v2" //Error

	cli := NewClient(cfg.Space, cfg.APIKey)
	_, err = cli.Delete(api, nil)
	if err == nil {
		t.Fatal(err)
	}
	s := fmt.Sprintf("%s", err)
	if s != "Error: StatusCode:400 Code:6 Message: Undefined resource. /api/v2/ MoreInfo:" {
		fmt.Println(err)
	}
}

func TestClient_Delete02(t *testing.T) {

	cfg, err := NewConfig()
	if err != nil {
		t.Fatal(err)
	}

	api := "api/v2/issues/0"
	apiKey := "xxxxxxxxxxxxxxx" //Error

	cli := NewClient(cfg.Space, apiKey)
	_, err = cli.Delete(api, nil)
	if err == nil {
		t.Fatal(err)
	}
	s := fmt.Sprintf("%s", err)
	if s != "Error: StatusCode:401 Code:11 Message: Authentication failure. MoreInfo:" {
		fmt.Println(err)
	}
}

package client

import (
	"io/ioutil"
	"os"
	"fmt"
	"path/filepath"
)

// Config is ...
type Config struct {
	Space  string
	APIKey string
}

// NewConfig is ...
func NewConfig() (*Config, error) {

	var err error
	cfg := new(Config)
	path := os.Getenv("BACKLOG_CLI")
	if path == "" {
		err := fmt.Errorf("Environment variable not specified. %s", "BACKLOG_CLI")
		return nil, err
	}
	cfg.Space, err = cfg.readfile(filepath.Join(path, "space.txt"))
	if err != nil {
		return nil, err
	}
	cfg.APIKey, err = cfg.readfile(filepath.Join(path, "apikey.txt"))
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func (cfg *Config) readfile(path string) (string, error) {
	// ファイル読み込み
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

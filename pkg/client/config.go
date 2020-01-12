package client

import (
	"io/ioutil"
	"os"
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
	cfg.Space, err = cfg.readfile(path + "/space.txt")
	if err != nil {
		return nil, err
	}
	cfg.APIKey, err = cfg.readfile(path + "/apikey.txt")
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
	return string(b), nil
}

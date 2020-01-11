package client

import (
	"io/ioutil"
	"os"
)

// Config is ...
type Config struct {
	space  string
	apiKey string
}

// NewConfig is ...
func NewConfig() (*Config, error) {

	var err error
	cfg := new(Config)
	path := os.Getenv("BACKLOG_CLI")
	cfg.space, err = cfg.readfile(path + "/space.txt")
	if err != nil {
		return nil, err
	}
	cfg.apiKey, err = cfg.readfile(path + "/apikey.txt")
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

// Setup is ...
func (cfg Config) Setup() (string, string) {
	return cfg.space, cfg.apiKey
}

func (cfg Config) readfile(path string) (string, error) {
	// ファイル読み込み
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	return string(b), nil
}

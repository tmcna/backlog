package backlogcli

import (
	"fmt"

	"github.com/tmcna/backlog/pkg/client"
)

// ProjectList function executes project list subcommand.
func ProjectList() error {
	cfg, err := client.NewConfig()
	if err != nil {
		err = fmt.Errorf("configuration error, %s", err)
		return err
	}
	p := client.NewProjects(cfg.Space, cfg.APIKey)
	r, err := p.List()
	if err != nil {
		return err
	}
	p.PrintCSV(r)
	return nil
}

// ProjectInfo function executes project info subcommand.
func ProjectInfo(projectKey string) error {
	cfg, err := client.NewConfig()
	if err != nil {
		err = fmt.Errorf("configuration error, %s", err)
		return err
	}
	p, err := client.NewProject(cfg.Space, cfg.APIKey, projectKey)
	if err != nil {
		return err
	}
	p.Print()
	return nil
}

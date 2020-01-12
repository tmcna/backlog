package backlogcli

import (
	"fmt"

	"github.com/tmcna/backlog/pkg/client"
)

// ActivityList function executes act subcommand.
func ActivityList() error {
	cfg, err := client.NewConfig()
	if err != nil {
		err = fmt.Errorf("configuration error, %s", err)
		return err
	}

	act := client.NewActivity(cfg.Space, cfg.APIKey, 32, client.DisplayOrderDesc)
	for act.HasNext() {
		r, err := act.List()
		if err != nil {
			return err
		}
		act.PrintCSV(r)
	}
	return nil
}

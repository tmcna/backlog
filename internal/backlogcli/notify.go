package backlogcli

import (
	"fmt"

	"github.com/tmcna/backlog/pkg/client"
)

// NotifyList function executes notify subcommand.
func NotifyList() error {
	cfg, err := client.NewConfig()
	if err != nil {
		err = fmt.Errorf("configuration error, %s", err)
		return err
	}
	n := client.NewNotification(cfg.Space, cfg.APIKey)
	r, err := n.List()
	if err != nil {
		return err
	}
	n.PrintCSV(r)
	return nil
}

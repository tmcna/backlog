package backlogcli

import (
	"fmt"

	"github.com/tmcna/backlog/pkg/client"
)

// SpaceUsage function executes space subcommand.
func SpaceUsage() error {
	cfg, err := client.NewConfig()
	if err != nil {
		err = fmt.Errorf("configuration error, %s", err)
		return err
	}
	sp := client.NewSpace(cfg.Space, cfg.APIKey)
	r, err := sp.GetSpaceUsage()
	if err != nil {
		return err
	}
	sp.PrintSpaceUsageCSV(r)
	return nil
}

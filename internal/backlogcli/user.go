package backlogcli

import (
	"fmt"

	"github.com/tmcna/backlog/pkg/client"
)

// UserList function user list subcommand.
func UserList() error {

	cfg, err := client.NewConfig()
	if err != nil {
		err = fmt.Errorf("configuration error, %s", err)
		return err
	}

	user := client.NewUser(cfg.Space, cfg.APIKey)
	r, err := user.List()
	if err != nil {
		return err
	}
	user.PrintCSV(r)
	return nil
}

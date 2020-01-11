package main

import (
	"os"

	command "github.com/tmcna/backlog/internal/backlogcli"
)

func main() {
	os.Exit(command.Cli(os.Args))
}

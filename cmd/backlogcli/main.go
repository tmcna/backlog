package main

import (
	"os"

	"github.com/tmcna/backlog/pkg/client"
	"github.com/tmcna/backlog/internal/backlogcli"
)

func main() {
	os.Exit(backlogcli.Cli(os.Args))
}

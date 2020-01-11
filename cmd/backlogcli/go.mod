module github.com/tmcna/backlog/cmd/backlogcli

go 1.13

replace (
	github.com/tmcna/backlog/internal/backlogcli => ../../internal/backlogcli
	github.com/tmcna/backlog/pkg/client => ../../pkg/client
)

require (
	github.com/BurntSushi/toml v0.3.1
	github.com/tealeg/xlsx v1.0.5
	github.com/urfave/cli/v2 v2.1.1
)

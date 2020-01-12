module github.com/tmcna/backlog

go 1.13

replace (
	github.com/tmcna/backlog/internal/backlogcli => ./internal/backlogcli
	github.com/tmcna/backlog/pkg/client => ./pkg/client
)

require (
	github.com/tmcna/backlog
	github.com/tmcna/backlog/pkg
	github.com/BurntSushi/toml v0.3.1
	github.com/tealeg/xlsx v1.0.5
	github.com/tmcna/backlog/internal/backlogcli v0.0.0-00010101000000-000000000000 // indirect
	github.com/tmcna/backlog/pkg/client v0.0.0-20200112034039-c5ea3606ec00 // indirect
	github.com/urfave/cli/v2 v2.1.1
)

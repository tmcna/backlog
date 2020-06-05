package backlogcli

import (
	"fmt"
	"os"

	cli "github.com/urfave/cli/v2"
)

// ExitCodeOK, ExitCodeError is return code.
const (
	ExitCodeOK    int = iota // 0
	ExitCodeError            // 1
)

var debug bool = false

// Cli function parses the command line and then executes application.
func Cli(args []string) int {

	app := &cli.App{
		Name:    "backlogcli",
		Version: Version,
		Usage:   "A CLI application for Backlog users.",
		Commands: []*cli.Command{
			{
				Name:    "user",
				Aliases: []string{"u"},
				Usage:   "List of users in your space.",
				Subcommands: []*cli.Command{
					{
						Name:  "list",
						Aliases: []string{"ls"},
						Usage: "List of users.",
						Action: func(c *cli.Context) error {
								if err := UserList(); err != nil {
									return err
								}
								return nil
							},
					},
				},
			},
			{
				Name:    "act",
				Aliases: []string{"a"},
				Usage:   "Recent updates in your space.",
				Action: func(c *cli.Context) error {
					if err := ActivityList(); err != nil {
						return err
					}
					return nil
				},
			},
			{
				Name:    "notify",
				Aliases: []string{"n"},
				Usage:   "Updates space notification.",
				Action: func(c *cli.Context) error {
					if err := NotifyList(); err != nil {
						return err
					}
					return nil
				},
			},
			{
				Name:    "space",
				Aliases: []string{"s"},
				Usage:   "Information about space disk usage.",
				Action: func(c *cli.Context) error {
					if err := SpaceUsage(); err != nil {
						return err
					}
					return nil
				},
			},
			{
				Name:    "project",
				Aliases: []string{"p"},
				Usage:   "Operations project",
				Subcommands: []*cli.Command{
					{
						Name:  "list",
						Aliases: []string{"ls"},
						Usage: "List of projects.",
						Action: func(c *cli.Context) error {
							if err := ProjectList(); err != nil {
								return err
							}
							return nil
						},
					},
					{
						Name:  "info",
						Usage: "print information about the project.",
						Action: func(c *cli.Context) error {
							projectKey := c.Args().First()
							if projectKey == "" {
								err := fmt.Errorf("Error: Argument not found. %s", "Project key")
								return err
							}
							if err := ProjectInfo(projectKey); err != nil {
								return err
							}
							return nil
						},
					},
				},
			},
			{
				Name:    "issue",
				Aliases: []string{"i"},
				Usage:   "Operations issue",
				Subcommands: []*cli.Command{
					{
						Name:  "list",
						Aliases: []string{"ls"},
						Usage: "List of issues.",
						Action: func(c *cli.Context) error {
							if err := IssueList(); err != nil {
								return err
							}
							return nil
						},
					},
					{
						Name:  "info",
						Usage: "Information of issues.",
						Action: func(c *cli.Context) error {
							if err := IssueInfo(c.Args().First()); err != nil {
								return err
							}
							return nil
						},
					},
					{
						Name:  "add",
						Usage: "Add issue.",
						Action: func(c *cli.Context) error {
							file := c.Args().First()
							if err := IssueAdd(file); err != nil {
								return err
							}
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "Update issue.",
						Aliases: []string{"u"},
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "file", Aliases: []string{"f"}},
							&cli.StringFlag{Name: "status", Aliases: []string{"s"}},
							&cli.StringFlag{Name: "assignee", Aliases: []string{"a"}},
							&cli.StringFlag{Name: "comment", Aliases: []string{"c"}},
							&cli.StringSliceFlag{Name: "notify", Aliases: []string{"n"}},
						},
						Action: func(c *cli.Context) error {
							file := c.String("file")
							status := c.String("status")
							assignee := c.String("assignee")
							comment := c.String("comment")
							//notifiedUsers := c.StringSlice("notify")
							//for i, n := range notifiedUsers {
							//	fmt.Printf("notify:%d %s\n", i, n)
							//}
							issueKey := c.Args().First()
							err := IssueUpdate(file, status, assignee, comment, issueKey)
							if err != nil {
								return err
							}
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "Deletes issue.",
						Action: func(c *cli.Context) error {
							if err := IssueDelete(c.Args().First()); err != nil {
								return err
							}
							return nil
						},
					},
				},
			},
			{
				Name:    "comment",
				Aliases: []string{"c"},
				Usage:   "Operations comment",
				Subcommands: []*cli.Command{
					{
						Name:  "add",
						Usage: "Add comment.",
						Action: func(c *cli.Context) error {
							if err := CommentAdd(c.Args().First(), c.Args().Get(1)); err != nil {
								return err
							}
							return nil
						},
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return ExitCodeError
	}
	return ExitCodeOK
}

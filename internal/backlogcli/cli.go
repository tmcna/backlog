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
				Action: func(c *cli.Context) error {
					s := NewSubcommand()
					err := s.UserList()
					if err != nil {
						return err
					}
					return nil
				},
			},
			{
				Name:    "act",
				Aliases: []string{"a"},
				Usage:   "Recent updates in your space.",
				Action: func(c *cli.Context) error {
					s := NewSubcommand()
					err := s.ActivityList()
					if err != nil {
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
					s := NewSubcommand()
					err := s.NotifyList()
					if err != nil {
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
					s := NewSubcommand()
					err := s.SpaceUsage()
					if err != nil {
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
						Usage: "List of projects.",
						Action: func(c *cli.Context) error {
							s := NewSubcommand()
							err := s.ProjectList()
							if err != nil {
								return err
							}
							return nil
						},
					},
					{
						Name:  "info",
						Usage: "print information about the project.",
						Action: func(c *cli.Context) error {
							s := NewSubcommand()
							projectKey := c.Args().First()
							if projectKey == "" {
								err := fmt.Errorf("Error: Argument not found. %s", "Project key")
								return err
							}
							err := s.ProjectInfo(projectKey)
							if err != nil {
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
						Usage: "List of issues.",
						Action: func(c *cli.Context) error {
							s := NewSubcommand()
							err := s.IssueList()
							if err != nil {
								return err
							}
							return nil
						},
					},
					{
						Name:  "add",
						Usage: "Add issue.",
						Action: func(c *cli.Context) error {
							s := NewSubcommand()
							file := c.Args().First()
							err := s.IssueAdd(file)
							if err != nil {
								return err
							}
							return nil
						},
					},
					{
						Name:  "update",
						Usage: "Update issue.",
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
							s := NewSubcommand()
							issueKey := c.Args().First()
							err := s.IssueUpdate(file, status, assignee, comment, issueKey)
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
							s := NewSubcommand()
							err := s.IssueDelete(c.Args().First())
							if err != nil {
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
							s := NewSubcommand()
							err := s.CommentAdd(c.Args().First(), c.Args().Get(1))
							if err != nil {
								return err
							}
							return nil
						},
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return ExitCodeError
	}
	return ExitCodeOK
}

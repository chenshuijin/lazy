package main

import cli "gopkg.in/urfave/cli.v1"

var (
	newProjectCmd = cli.Command{
		Name:        "new",
		Category:    "New Project",
		Description: "Create new project",
		Usage:       "Create new project",
		Subcommands: []cli.Command{
			newClientCmd,
			newServiceCmd,
		},
	}
	newClientCmd = cli.Command{
		Action:      newCliProject,
		Name:        "cli",
		Category:    "New Project",
		Description: "Create new client project",
		Usage:       "Create new client project",
		Flags: []cli.Flag{
			nameFlag,
			descFlag,
		},
	}
	newServiceCmd = cli.Command{
		Action:      newServiceProject,
		Name:        "svr",
		Category:    "New Project",
		Description: "Create new service project",
		Usage:       "Create new service project",
		Flags: []cli.Flag{
			nameFlag,
			descFlag,
		},
	}
)

package main

import cli "gopkg.in/urfave/cli.v1"

var (
	projectName string
	desc        string

	nameFlag = cli.StringFlag{
		Name:        "n",
		Usage:       "Project name, always the path in the GOPATH",
		Value:       "golazysample",
		Destination: &projectName,
	}

	descFlag = cli.StringFlag{
		Name:        "d",
		Usage:       "Description of the project",
		Value:       "",
		Destination: &desc,
	}
)

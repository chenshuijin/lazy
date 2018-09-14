package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"

	cli "gopkg.in/urfave/cli.v1"
)

var (
	version   string
	commit    string
	branch    string
	compileAt string
	config    string
)

func main() {
	cli.AppHelpTemplate = appHelpTemplate
	app := cli.NewApp()

	app.Action = nil
	app.Name = "golazy"
	timestamp, _ := strconv.ParseInt(compileAt, 10, 64)
	app.Compiled = time.Unix(timestamp, 0)
	app.Version = fmt.Sprintf("%s\n branch: %s\n commit: %s\n compileAt: %s", version, branch, commit, app.Compiled)

	app.Usage = "The golazy command line interface"
	app.Copyright = "Copyright 2018-2019 The Authors chenshuijin<785795635@qq.com>"

	app.Flags = append(app.Flags)

	sort.Sort(cli.FlagsByName(app.Flags))

	app.Commands = []cli.Command{
		newProjectCmd,
	}
	sort.Sort(cli.CommandsByName(app.Commands))

	app.Run(os.Args)
}

func golazy(ctx *cli.Context) error {
	return nil
}

package generator

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// CreateMainFile create a go main file
func CreateMainFile(fileName, binName string) {
	createDir(fileName)
	mainf := strings.Replace(maintpl, "[[bin]]", binName, -1)
	if err := ioutil.WriteFile(fileName, []byte(mainf), 0644); err != nil {
		fmt.Printf("write file [%s] failed:%v\n", fileName, err)
	}
}

// CreateTestFile create a go test file
func CreateTestFile(fileName string) {
	if !strings.HasSuffix(fileName, "_test.go") {
		return
	}
	createDir(fileName)
	if err := ioutil.WriteFile(fileName, []byte(mainttpl), 0644); err != nil {
		fmt.Printf("write file [%s] failed:%v\n", fileName, err)
	}
}

var mainttpl = `package main

import "testing"

func TestM(t *testing.T) {

}

`

var maintpl = `package main

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
	app := cli.NewApp()
	app.Action = [[bin]]
	app.Name = "[[bin]]"
	timestamp, _ := strconv.ParseInt(compileAt, 10, 64)
	app.Compiled = time.Unix(timestamp, 0)
	app.Version = fmt.Sprintf("%s\n branch: %s\n commit: %s\n compileAt:%s", version, branch, commit, app.Compiled)

	app.Usage = "The [[bin]] command line interface"
	app.Copyright = "Copyright 2017-2018 The Authors"

	sort.Sort(cli.FlagsByName(app.Flags))

	app.Commands = []cli.Command{}
	sort.Sort(cli.CommandsByName(app.Commands))

	app.Run(os.Args)
}

func [[bin]](ctx *cli.Context) error {
	return nil
}

`

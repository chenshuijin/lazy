package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/go-ray/lazy/generator"
	cli "gopkg.in/urfave/cli.v1"
)

func newCliProject(ctx *cli.Context) error {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath, _ = os.Getwd()
	}
	projectRoot := filepath.Join(gopath, "src", projectName)
	projectRoot, err := filepath.Abs(projectRoot)
	if err != nil {
		fmt.Printf("create dir [%s] failed:%v\n", projectRoot, err)
		return err
	}
	_, binName := filepath.Split(projectRoot)

	mainfile := filepath.Join(projectRoot, "cmd", binName, "main.go")
	generator.CreateMainFile(mainfile, binName)

	mainttfile := filepath.Join(projectRoot, "cmd", binName, "main_test.go")
	generator.CreateTestFile(mainttfile)

	makefile := filepath.Join(projectRoot, "makefile")
	generator.CreateMakeFileWithoutContrib(makefile, binName)

	fmt.Println("binName:", binName)
	fmt.Println("gopath:", projectRoot)

	return nil
}

func newServiceProject(ctx *cli.Context) error {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath, _ = os.Getwd()
	}
	projectRoot := filepath.Join(gopath, "src", projectName)
	projectRoot, err := filepath.Abs(projectRoot)
	if err != nil {
		fmt.Printf("create dir [%s] failed:%v\n", projectRoot, err)
		return err
	}
	_, binName := filepath.Split(projectRoot)

	mainfile := filepath.Join(projectRoot, "cmd", binName, "main.go")
	generator.CreateMainFile(mainfile, binName)

	mainttfile := filepath.Join(projectRoot, "cmd", binName, "main_test.go")
	generator.CreateTestFile(mainttfile)

	makefile := filepath.Join(projectRoot, "makefile")
	generator.CreateMakeFile(makefile, binName)

	installfile := filepath.Join(projectRoot, "contrib", "script", "install.sh")
	generator.CreateInstallFile(installfile, binName)

	servicefile := filepath.Join(projectRoot, "contrib", "systemd", binName+".service")
	generator.CreateServiceFile(servicefile, binName, desc)

	fmt.Println("binName:", binName)
	fmt.Println("gopath:", projectRoot)
	return nil
}

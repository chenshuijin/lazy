package generator

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// CreateMakeFile create make file
func CreateMakeFile(fileName, binName string) error {
	createDir(fileName)
	mkfile := strings.Replace(mkfiletpl, "[[bin]]", binName, -1)
	mkfile = strings.Replace(mkfile, "[[u]]", binName, -1)
	if err := ioutil.WriteFile(fileName, []byte(mkfile), 0644); err != nil {
		fmt.Printf("write file [%s] failed:%v\n", fileName, err)
		return err
	}
	return nil
}

// CreateMakeFileWithoutContrib create make file without contrib files
func CreateMakeFileWithoutContrib(fileName, binName string) error {
	createDir(fileName)
	mtp := mktpl
	pks := mtp.Targets[8].Cmd
	mtp.Targets[8].Cmd = pks[1:]
	mtp.Params[8].Value = ""
	mkfile := strings.Replace(mtp.String(), "[[bin]]", binName, -1)
	mkfile = strings.Replace(mkfile, "[[u]]", binName, -1)
	if err := ioutil.WriteFile(fileName, []byte(mkfile), 0644); err != nil {
		fmt.Printf("write file [%s] failed:%v\n", fileName, err)
		return err
	}
	return nil
}

// GetMake just for test
func GetMake() string {
	fmt.Println(mktpl.String())
	return ""
}

var mktpl = Make{
	Params: []Param{
		{"VERSION?", "0.0.1", nil},
		{"COMMIT", "$(shell git rev-parse HEAD)", nil},
		{"BRANCH", "$(shell git rev-parse --abbrev-ref HEAD)", nil},
		{"CURRENT_DIR", "$(shell pwd)", nil},
		{"BUILD_DIR", "${CURRENT_DIR}", nil},
		{"BINARY", "[[bin]]", nil},
		{"TARFILE", "${BINARY}.tar.bz2", nil},
		{"CONTRIBFILES", "contrib/systemd/[[bin]].service contrib/script/install.sh", nil},
		{"RMBIN", "${BINARY}.service install.sh", nil},
		{"VET_REPORT", "vet.report", nil},
		{"LINT_REPORT", "lint.report", nil},
		{"TEST_REPORT", "test.report", nil},
		{"TEST_XUNIT_REPORT", "test.report.xml", nil},
		{"OUTPUT", "${CURRENT_DIR}/bin", nil},
		{
			"LDFLAGS",
			"-ldflags \"-X main.version=${VERSION} -X main.commit=${COMMIT} -X main.branch=${BRANCH} -X main.compileAt=`date +%s`\"",
			[]string{"Setup the -ldflags option for go build here, interpolate the variable values"},
		},
	},
	Targets: []Target{
		{".PHONY", []string{"build", "clean", "dep", "lint", "run", "test", "vet", "package"}, nil, []string{"Build the project"}},
		{"all", []string{"clean", "vet", "fmt", "lint", "build", "test", "package"}, nil, nil},
		{"dep", nil, []string{"dep ensure -v"}, nil},
		{"build", nil, []string{"cd cmd/$(BINARY); go build $(LDFLAGS) -o $(OUTPUT)/$(BINARY)"}, nil},
		{"test", nil, []string{"env GOCACHE=off go test ./... 2>&1 | tee $(TEST_REPORT); go2xunit -fail -input $(TEST_REPORT) -output $(TEST_XUNIT_REPORT)"}, nil},
		{"vet", nil, []string{"go vet $$(go list ./...) 2>&1 | tee $(VET_REPORT)"}, nil},
		{"fmt", nil, []string{"goimports -w $$(go list -f \"{{.Dir}}\" ./... | grep -v /vendor/)"}, nil},
		{"lint", nil, []string{"#golint $$(go list ./...) | sed 's:^$(BUILD_DIR)/::' | tee $(LINT_REPORT)"}, nil},
		{"package", nil, []string{"for i in ${CONTRIBFILES}; do cp $$i $(OUTPUT); done", "cd $(OUTPUT) && tar -jcf $(TARFILE) $(BINARY) ${RMBIN}", "cd $(OUTPUT) && for i in ${RMBIN}; do rm -f $$i; done"}, nil},
		{"clean", nil, []string{"-rm -f $(VET_REPORT)", "-rm -f $(LINT_REPORT)", "-rm -f $(TEST_REPORT)", "-rm -f $(TEST_XUNIT_REPORT)", "-rm -f $(BINARY)", "-rm -f $(OUTPUT)"}, nil},
	},
}

var mkfiletpl = `VERSION?=0.0.1

COMMIT=$(shell git rev-parse HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)

CURRENT_DIR=$(shell pwd)
BUILD_DIR=${CURRENT_DIR}
BINARY=[[bin]]
TARFILE=${BINARY}.tar.bz2
CONTRIBFILES=contrib/systemd/[[bin]].service contrib/script/install.sh
RMBIN=${BINARY}.service install.sh

VET_REPORT=vet.report
LINT_REPORT=lint.report
TEST_REPORT=test.report
TEST_XUNIT_REPORT=test.report.xml

OUTPUT=${CURRENT_DIR}/bin

OS := $(shell uname -s)
ifeq ($(OS),Darwin)

else

endif

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS = -ldflags "-X main.version=${VERSION} -X main.commit=${COMMIT} -X main.branch=${BRANCH} -X main.compileAt=` + "`date +%s`\"" + `

# Build the project
.PHONY: build clean dep lint run test vet package

all: clean vet fmt lint build test package

dep:
	dep ensure -v

build:
	cd cmd/$(BINARY); go build $(LDFLAGS) -o $(OUTPUT)/$(BINARY)

test:
	env GOCACHE=off go test ./... 2>&1 | tee $(TEST_REPORT); go2xunit -fail -input $(TEST_REPORT) -output $(TEST_XUNIT_REPORT)

vet:
	go vet $$(go list ./...) 2>&1 | tee $(VET_REPORT)

fmt:
	goimports -w $$(go list -f "{{.Dir}}" ./... | grep -v /vendor/)

lint:
#	golint $$(go list ./...) | sed "s:^$(BUILD_DIR)/::" | tee $(LINT_REPORT)

package:
	for i in ${CONTRIBFILES}; do cp $$i $(OUTPUT); done
	cd $(OUTPUT) && tar -jcf $(TARFILE) $(BINARY) ${RMBIN}
	cd $(OUTPUT) && for i in ${RMBIN}; do rm -f $$i; done

clean:
	-rm -f $(VET_REPORT)
	-rm -f $(LINT_REPORT)
	-rm -f $(TEST_REPORT)
	-rm -f $(TEST_XUNIT_REPORT)
	-rm -f $(BINARY)
	-rm -f $(OUTPUT)/*
`

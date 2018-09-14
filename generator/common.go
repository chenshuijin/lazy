package generator

import (
	"fmt"
	"os"
	"path/filepath"
)

func createDir(p string) {
	p = filepath.Clean(p)
	d, _ := filepath.Split(p)
	if err := os.MkdirAll(d, os.ModePerm); err != nil {
		fmt.Printf("create dir [%s] failed:%v\n", p, err)
	}
}

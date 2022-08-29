package gosh

import (
	"fmt"
	"os"
	"path/filepath"
)

func Chdir(path string) {
	currPath, err := filepath.Abs(path)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	err = os.Chdir(currPath)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}

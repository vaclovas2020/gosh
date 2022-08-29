package gosh

import (
	"errors"
	"os"
	"os/exec"
	"strings"
)

func ExecDir(input string, currPath string, path string) (bool, error) {
	if _, err := os.Stat(path + strings.Split(input, " ")[0]); errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else {
		arg := []string{}
		if len(strings.Split(input, " ")) > 1 {
			arg = strings.Split(input, " ")[1:]
		}
		cmd := exec.Command(path+strings.Split(input, " ")[0], arg...)
		cmd.Dir = currPath
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		err := cmd.Start()
		if err != nil {
			return true, err
		}
		err = cmd.Wait()
		if err != nil {
			return true, err
		}
		return true, nil
	}
}

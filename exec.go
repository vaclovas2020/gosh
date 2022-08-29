package gosh

import (
	"fmt"
	"os"
)

func Exec(input string) {
	currPath, err := os.Getwd()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	success, err := ExecDir(input, currPath, "/sbin/")
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	if !success {
		success, err = ExecDir(input, currPath, "/bin/")
		if err != nil {
			fmt.Printf("%v\n", err)
		}
		if !success {
			success, err = ExecDir(input, currPath, "/usr/sbin/")
			if err != nil {
				fmt.Printf("%v\n", err)
			}
			if !success {
				success, err = ExecDir(input, currPath, "/usr/bin/")
				if err != nil {
					fmt.Printf("%v\n", err)
				}
				if !success {
					success, err = ExecDir(input, currPath, currPath+"/")
					if err != nil {
						fmt.Printf("%v\n", err)
					}
					if !success {
						NotFound(input)
					}
				}
			}
		}
	}
}

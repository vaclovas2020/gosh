package gosh

import "fmt"

func NotFound(command string) {
	fmt.Printf("Command `%s` not found..\n", command)
}

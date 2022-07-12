package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/inancgumus/screen"
	"webimizer.dev/gosh"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	screen.Clear()
	fmt.Println("Welcome to Gosh v0.0.1 (Linux based Go shell)")
	for {
		fmt.Printf("%s$ ", os.Args[0])
		input, err := buf.ReadString('\n')
		if err != nil {
			panic(err)
		}
		input = strings.ReplaceAll(input, "\n", "")
		switch input {
		case "help":
			gosh.Help()
		case "restart":
			syscall.Reboot(syscall.LINUX_REBOOT_CMD_RESTART)
		case "shutdown":
			syscall.Reboot(syscall.LINUX_REBOOT_CMD_POWER_OFF)
		default:
			gosh.NotFound(input)
		}
	}
}

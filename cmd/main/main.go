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
	screen.MoveTopLeft()
	fmt.Println("Welcome to Gosh v0.0.1 (Linux based Go shell)")
	fmt.Println()
	fmt.Println("Copyright (c) 2022 Vaclovas Lapinskis")
	fmt.Println()
	for {
		fmt.Printf("%s$ ", os.Args[0])
		input, err := buf.ReadString('\n')
		if err != nil {
			panic(err)
		}
		input = strings.ReplaceAll(input, "\n", "")
		if input == "" {
			continue
		}
		switch input {
		case "clear":
			screen.Clear()
			screen.MoveTopLeft()
		case "help":
			gosh.Help()
		case "restart":
			syscall.Reboot(syscall.LINUX_REBOOT_CMD_RESTART)
		case "shutdown":
			syscall.Reboot(syscall.LINUX_REBOOT_CMD_POWER_OFF)
		case "goroutines_test":
			gosh.GoRoutinesTest()
		default:
			gosh.NotFound(input)
		}
	}
}

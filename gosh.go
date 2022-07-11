package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	buf := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s$ ", os.Args[0])
		input, err := buf.ReadString('\n')
		if err != nil {
			panic(err)
		}
		input = strings.ReplaceAll(input, "\n", "")
		switch input {
		case "hello":
			hello()
		case "exit":
			os.Exit(0)
		default:
			notFound(input)
		}
	}
}

func hello() {
	fmt.Println("Hello world")
}

func notFound(command string) {
	fmt.Printf("Command `%s` not found..\n", command)
}

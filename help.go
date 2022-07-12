package gosh

import "fmt"

func Help() {
	fmt.Println(`Commands:
	help:
		help screen
	shutdown:
		power off computer
	restart:
		restart computer
	`)
}

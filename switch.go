package gosh

import (
	"strings"

	"github.com/inancgumus/screen"
)

/* Execute command from input */
func Switch(input string) {
	switch strings.Split(input, " ")[0] {
	case "cd":
		if len(strings.Split(input, " ")) > 1 {
			Chdir(strings.Split(input, " ")[1])
		}
	case "clear":
		screen.Clear()
		screen.MoveTopLeft()
	default:
		Exec(input)
	}
}

package gosh

import (
	"bufio"
	"os"
	"strings"
)

/* Init gosh shell */
func InitGosh() {
	InitScreen()
	buf := bufio.NewReader(os.Stdin)
	for {
		Wd()
		input := ReadBuf(buf)
		if input == "" {
			continue
		}
		commands := strings.Split(input, "&&")
		for _, s := range commands {
			Switch(strings.Trim(s, " "))
		}
	}
}

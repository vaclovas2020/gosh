package gosh

import (
	"bufio"
	"fmt"
	"strings"
)

/* Read stdin buffer and return string */
func ReadBuf(buf *bufio.Reader) string {
	input, err := buf.ReadString('\n')
	if err != nil {
		fmt.Printf("%v\n", err)
		return ""
	}
	return strings.ReplaceAll(input, "\n", "")
}

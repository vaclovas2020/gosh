package gosh

import (
	"fmt"
	"os"
)

/* Set current work dir */
func Wd() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("%s $ ", wd)
}

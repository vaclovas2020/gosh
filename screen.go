package gosh

import (
	"fmt"
	"runtime"

	"github.com/inancgumus/screen"
)

/* Init start screen */
func InitScreen() {
	screen.Clear()
	screen.MoveTopLeft()
	fmt.Println("Welcome to Gosh v0.0.2 (Linux based Go shell)")
	fmt.Println()
	fmt.Println("Linux kernel version:", GetUname())
	fmt.Println("Go runtime version:", runtime.Version())
	fmt.Println("The number of CPU Cores:", runtime.NumCPU())
	fmt.Println()
	fmt.Println("Copyright (c) 2022 Vaclovas Lapinskis")
	fmt.Println()
}

package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"syscall"
	"time"
)

func main() {
	var list bool
	flag.BoolVar(&list, "l", false, "print as detail list")
	flag.Parse()
	path := flag.Arg(0)
	if path == "" {
		path = "./"
	}
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	for _, f := range files {
		if list {
			printListEntry(f)
		} else {
			printEntry(f)
		}
	}
	fmt.Printf("\n")
}

func printEntry(f fs.DirEntry) {
	if f.IsDir() {
		fmt.Printf("\033[;34m%s\033[0m ", f.Name())
	} else {
		fmt.Printf("\033[;32m%s\033[0m ", f.Name())
	}
}

func printListEntry(f fs.DirEntry) {
	info, err := f.Info()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(2)
	}
	var GID, UID int
	if stat, ok := info.Sys().(*syscall.Stat_t); ok {
		UID = int(stat.Uid)
		GID = int(stat.Gid)
	}
	if f.IsDir() {
		fmt.Printf("%s %d %d %8d %s \033[;34m%s\033[0m\n", info.Mode().String(), UID, GID, info.Size(), info.ModTime().Format(time.UnixDate), f.Name())
	} else {
		fmt.Printf("%s %d %d %8d %s \033[;32m%s\033[0m\n", info.Mode().String(), UID, GID, info.Size(), info.ModTime().Format(time.UnixDate), f.Name())
	}
}

package main

import "syscall"

func main() {
	syscall.Reboot(syscall.LINUX_REBOOT_CMD_RESTART)
}

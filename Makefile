build:
	go build -o bin/gosh gosh.go
	gcc -ffreestanding -no-pie init/init.c -o sbin/init
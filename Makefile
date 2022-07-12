build:
	go build -o bin/gosh gosh.go
	gcc -nostdlib -ffreestanding -no-pie ./init/* -o sbin/init
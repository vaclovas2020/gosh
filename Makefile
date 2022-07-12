build:
	go build -o bin/gosh ./cmd/main/main.go
	gcc -nostdlib -ffreestanding -no-pie ./init/* -o sbin/init
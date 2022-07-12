build:
	go build -o bin/gosh gosh.go
	gcc -static init/init.c -o sbin/init
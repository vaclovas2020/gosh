build:
	go build -o ./_install/bin/gosh ./cmd/gosh/main.go
	go build -o ./_install/bin/ls ./cmd/ls/main.go
	go build -o ./_install/sbin/restart ./cmd/sbin/restart/main.go
	go build -o ./_install/sbin/shutdown ./cmd/sbin/shutdown/main.go
	gcc -nostdlib -ffreestanding -no-pie ./init/* -o ./_install/init
install:
	cd  ./_install/ && find . -print0 | cpio --null -ov --format=newc -R 0:0 | gzip -9 > ../initramfs.cpio.gz
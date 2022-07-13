package gosh

import "fmt"

func GoRoutinesTest() {
	count := 4
	buff := make(chan string, count)
	for i := 0; i < count; i++ {
		go increment(i, buff)
	}
	go printWorker(buff, count)
}

func increment(in int, buff chan<- string) {
	buff <- fmt.Sprintf("%v", in)
}

func printWorker(buff <-chan string, count int) {
	for i := 0; i < count; i++ {
		fmt.Println(<-buff)
	}
}

package gosh

import "fmt"

func GoRoutinesTest() {
	fmt.Println()
	count := 4
	buff := make(chan string, count)
	done := make(chan bool, 1)
	for i := 0; i < count; i++ {
		go increment(i, buff)
	}
	go printWorker(buff, count, done)
	<-done
}

func increment(in int, buff chan<- string) {
	buff <- fmt.Sprintf("%v", in)
}

func printWorker(buff <-chan string, count int, done chan<- bool) {
	for i := 0; i < count; i++ {
		fmt.Println(<-buff)
	}
	done <- true
}

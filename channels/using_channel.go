package main

import "fmt"

func main() {

	ch := make(chan int)
	// Send
	go foo(ch)

	// Receive
	bar(ch)

	fmt.Println("exiting.. ")
}

// Send
func foo(ch chan<- int) {
	ch <- 9090
}

// Receive
func bar(ch <-chan int) {
	fmt.Println(<-ch)
}

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
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}

// Receive
func bar(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}

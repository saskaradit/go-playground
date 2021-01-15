package main

import "fmt"

func main() {
	// BELOW DOES NOT RUNT //
	// ch := make(chan int)

	// ch <- 42

	// fmt.Println(<-ch)

	ch := make(chan int)

	go func() {
		ch <- 42
	}()

	fmt.Println(<-ch)
}

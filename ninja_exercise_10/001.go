package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		ch <- 9090
	}()

	fmt.Println(<-ch)
}

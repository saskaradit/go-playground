package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		ch <- 9090
	}()
	fmt.Println(<-ch)

	fmt.Println("-----")
	fmt.Printf("%T\n", ch)
}

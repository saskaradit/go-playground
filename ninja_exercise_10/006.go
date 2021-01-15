package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
		close(ch)
	}()
	func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()
	fmt.Println("exiting")

}

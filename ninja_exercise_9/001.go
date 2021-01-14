package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("early", runtime.NumCPU())
	fmt.Println("early", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("Hellow")
		wg.Done()
	}()
	go func() {
		fmt.Println("Worldds")
		wg.Done()
	}()

	fmt.Println("mid", runtime.NumCPU())
	fmt.Println("mid", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("end", runtime.NumCPU())
	fmt.Println("end", runtime.NumGoroutine())

	fmt.Println("exiting...")
}

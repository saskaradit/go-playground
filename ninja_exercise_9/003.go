package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPU", runtime.NumCPU())
	fmt.Println("Go Routine:", runtime.NumGoroutine())
	counter := 0

	const gs = 20
	var wg sync.WaitGroup
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			runtime.Gosched()
			counter++
			wg.Done()
		}()
		fmt.Println("Go Routine:", runtime.NumGoroutine())
	}
	wg.Wait()

}

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPU", runtime.NumCPU())
	fmt.Println("Go Routine before:", runtime.NumGoroutine())
	counter := 0

	const gs = 20
	var wg sync.WaitGroup
	var mute sync.Mutex
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			mute.Lock()
			runtime.Gosched()
			counter++
			mute.Unlock()
			wg.Done()
		}()
		fmt.Println("Go Routine after:", runtime.NumGoroutine())
	}
	wg.Wait()

}

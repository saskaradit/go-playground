// Solution for race condition

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

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	var mute sync.Mutex //mutex to prevent race condition

	for i := 0; i < gs; i++ {
		go func() {
			mute.Lock()
			v := counter
			// time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			mute.Unlock()
			wg.Done()
		}()
		fmt.Println("Go Routine:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Go Routine after:", runtime.NumGoroutine())
	fmt.Println("Counter", counter)
}

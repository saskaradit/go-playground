package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPU", runtime.NumCPU())
	fmt.Println("Go Routine:", runtime.NumGoroutine())
	var counter int64

	const gs = 20
	var wg sync.WaitGroup
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("atomic", atomic.LoadInt64(&counter))
			wg.Done()
		}()
		fmt.Println("Go Routine:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println(counter)

}

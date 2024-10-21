package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	const loopCount = 100
	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println("GoRoutines: ", runtime.NumGoroutine())
	var wg sync.WaitGroup
	wg.Add(loopCount)
	var count int64
	for i := 0; i < loopCount; i++ {
		go func() {
			atomic.AddInt64(&count, 1)
			runtime.Gosched()
			wg.Done()
		}()
		fmt.Println("GoRoutines: ", runtime.NumGoroutine())
	}

	fmt.Println("GoRoutines: ", runtime.NumGoroutine())
	fmt.Println("Count: ", count)

	wg.Wait()

}

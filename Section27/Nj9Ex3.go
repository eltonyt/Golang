package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	const loopCount = 100
	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println("GoRoutines: ", runtime.NumGoroutine())
	var wg sync.WaitGroup
	wg.Add(loopCount)
	count := 0
	for i := 0; i < loopCount; i++ {
		go func() {
			v := count
			v++
			runtime.Gosched()
			count = v
			wg.Done()
		}()
		fmt.Println("GoRoutines: ", runtime.NumGoroutine())
	}

	fmt.Println("GoRoutines: ", runtime.NumGoroutine())
	fmt.Println("Count: ", count)

	wg.Wait()

}

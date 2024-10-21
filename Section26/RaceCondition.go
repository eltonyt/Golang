package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println("GoRoutines: ", runtime.NumGoroutine())
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	counter := 0
	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			// time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("GoRoutines: ", runtime.NumGoroutine())
	}

	fmt.Println("GoRoutines: ", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("Counter: ", counter)

}

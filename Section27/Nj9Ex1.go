package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println("GoRoutines: ", runtime.NumGoroutine())
	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("GoRoutines: ", runtime.NumGoroutine())
	go t1(&wg)
	fmt.Println("GoRoutines: ", runtime.NumGoroutine())
	go t2(&wg)
	fmt.Println("GoRoutines: ", runtime.NumGoroutine())
	wg.Wait()

}

func t1(wg *sync.WaitGroup) {
	fmt.Println("Hello from function t1")
	wg.Done()
}

func t2(wg *sync.WaitGroup) {
	fmt.Println("Hello from function t2")
	wg.Done()
}

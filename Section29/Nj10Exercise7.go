package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	const countRoutine = 10
	wg.Add(countRoutine)
	for i := range countRoutine {
		go createChannel(i, &wg)
	}
	wg.Wait()
}

func createChannel(i int, wg *sync.WaitGroup) {
	c := make(chan int)
	go func() {
		for i := range 10 {
			c <- i
		}
		close(c)
	}()

	for v := range c {
		fmt.Println(i, v)

	}
	wg.Done()
}

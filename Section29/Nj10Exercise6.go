package main

import "fmt"

func main() {
	c := make(chan int, 100)
	for i := range 100 {
		c <- i
	}
	close(c)

	for v := range c {
		fmt.Println(v)
	}
}

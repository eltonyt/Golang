package main

import "fmt"

func main() {
	q := make(chan int)
	c := gen1(q)

	receive2(c, q)
	for v := range q {
		fmt.Println(v)
	}

	fmt.Println("about to exit")

}

func receive2(c <-chan int, q chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q:
			return
		}
	}
}

func gen1(q chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 1
		close(c)
		close(q)
	}()
	return c
}

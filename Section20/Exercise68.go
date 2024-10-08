package main

import "fmt"

func main() {
	x := func() int {
		return 1
	}()
	fmt.Printf("Value of x is %d", x)
}

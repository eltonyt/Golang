package main

import "fmt"

func times2() func() int {
	x := 1
	return func() int {
		x *= 2
		return x
	}
}

func main() {
	x := times2()
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
}

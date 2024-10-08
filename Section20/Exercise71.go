package main

import "fmt"

func operation(i int, y int, f func(int, int) int) int {
	return f(i, y)
}

func add(i int, y int) int {
	return i + y
}

func main() {
	x := operation(1, 2, add)
	fmt.Printf("Outcome after the operation is: %d", x)
}

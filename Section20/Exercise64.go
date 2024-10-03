package main

import "fmt"

func main() {
	fmt.Printf("Substract of %d and %d is %d", 5, 5, 0)
}

func substract(a int, b int) int {
	return a - b
}

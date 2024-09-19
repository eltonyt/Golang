package main

import "fmt"

func main() {
	xi := [5]int{1, 2, 3, 4, 5}
	for i, v := range xi {
		fmt.Printf("Value at index %d is %v \n", i, v)
	}
}

package main

import "fmt"

func main() {
	xi := []int{}
	xi = append(xi, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51)
	for i, v := range xi {
		fmt.Printf("Type at Index %d is %T, Value i %v \n", i, v, v)
	}
}

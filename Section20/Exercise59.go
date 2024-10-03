package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum1 := sum(x...)
	sum2 := sum2(x)
	fmt.Printf("Sum1 is %d ", sum1)
	fmt.Printf("Sum2 is %d ", sum2)
}

func sum(x ...int) int {
	result := 0
	for _, v := range x {
		result += v
	}
	return result
}

func sum2(x []int) int {
	result := 0
	for _, v := range x {
		result += v
	}
	return result
}

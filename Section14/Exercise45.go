package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Printf("Initialized Slice: %v \n", x)

	x = append(x, 53, 54, 55)
	fmt.Printf("After appending 53, 54, and 54 to the slice, new slice is: %v \n", x)

	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Printf("After appending elements from y to x, new slice is: %v \n", x)

}

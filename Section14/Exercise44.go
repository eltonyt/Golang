package main

import "fmt"

func main() {
	xi := []int{}
	xi = append(xi, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51)
	x1 := xi[0:5]
	x2 := xi[5:]
	x3 := xi[2:7]
	x4 := xi[1:6]

	fmt.Printf("x1 is %v\n", x1)
	fmt.Printf("x2 is %v\n", x2)
	fmt.Printf("x3 is %v\n", x3)
	fmt.Printf("x4 is %v\n", x4)
}

package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	x = append(x[:3], x[6:]...)

	fmt.Printf("After deleting 45, 46, and 47 from the slice, the new slice is %v \n", x)

}

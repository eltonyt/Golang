package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := Intn(250)
	fmt.Printf("%v variable has value %d \n", "X", x)

	switch {
	case x >= 0 && x <= 100:
		fmt.Println("Between 0 and 100")
	case x >= 101 && x <= 200:
		fmt.Println("Between 100 and 200")
	case x >= 201 && x <= 250:
		fmt.Println("Between 201 and 250")
	default:
		fmt.Println("Bigger than 250")
	}
}

func Intn(n int) int {
	return rand.Intn(n)
}
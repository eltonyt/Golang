package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Int31n(10)
	var y int = rand.Intn(10)
	fmt.Printf("X has value %d.\n", x)
	fmt.Printf("Y has value %d. \n", y)
	switch {
	case x < 4 && y < 4:
		fmt.Println("X and Y are both less than 4.")
	case x > 6 && y > 6:
		fmt.Println("X and Y are both greater than 6.")
	case x >= 4 && x <= 6:
		fmt.Println("X is greater or equal to 4 and less than or equal to 6.")
	case y != 5:
		fmt.Println("Y is not 5.")
	default:
		fmt.Println("None of previous cases were met.")
	}
}

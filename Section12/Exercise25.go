package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 1; i <= 42; i++ {
		x := rand.Intn(5)
		switch x {
		case 0:
			fmt.Println("X is 0.")
		case 1:
			fmt.Println("X is 1.")
		case 2:
			fmt.Println("X is 2.")
		case 3:
			fmt.Println("X is 3.")
		case 4:
			fmt.Println("X is 4.")
		}
		fmt.Printf("This is Loop %d. \n", i)
	}

}

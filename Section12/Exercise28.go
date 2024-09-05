package main

import "fmt"

func main() {
	for i := range 1000 {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("Odd Number: %d \n", i)
	}
}

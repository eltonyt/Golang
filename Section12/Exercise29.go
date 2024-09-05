package main

import "fmt"

func main() {
	for i := range 5 {
		fmt.Printf("Outside Loop - %d\n", i)
		for j := range 5 {
			fmt.Printf("Inner Loop - %d\n", j)
		}
	}
}

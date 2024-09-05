package main

import "fmt"

func main() {
	var i = 0
	for i < 10 {
		fmt.Printf("Currently I is %d. \n", i)
		i++
		if i == 4 {
			fmt.Println("Jump out of the loop")
			break
		}
	}
}

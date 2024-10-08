package main

import "fmt"

func main() {
	x := func() string {
		return "Hello World"
	}
	y := x()
	fmt.Printf("Lenght of the word is: %d", len(y))
}

package main

import "fmt"

func main() {
	aRandomFunction()
}

func aRandomFunction() {
	defer fmt.Printf("Defer Print Function - This line should be at the bottom.")
	fmt.Println("Yo")
	fmt.Println("I'm running a function")
	fmt.Println("This is the end of the function")
}

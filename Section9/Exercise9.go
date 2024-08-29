package main

import "fmt"

var firstVariable = "1st global variable"

const secondVariable = "2nd global variable"

func main() {
	localVariable := "Local Variable"
	fmt.Printf("1st Global Variable: %v \n", firstVariable)
	fmt.Printf("2nd Global Variable: %s \n", secondVariable)
	fmt.Printf("Local Variable: %s \n", localVariable)
}

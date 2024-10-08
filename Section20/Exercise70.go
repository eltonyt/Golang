package main

import "fmt"

func outFunc() func() int {
	return func() int {
		return -1
	}
}

func main() {
	function := outFunc
	x := function()()
	fmt.Printf("The Value is %d", x)
}

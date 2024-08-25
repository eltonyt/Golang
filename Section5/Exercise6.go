package main

import "fmt"

func main() {
	var st string = "Hello World"
	var interger int = 333
	var f float32 = 222.33

	fmt.Printf("Value %v has type: %T \n", st, st)
	fmt.Printf("Value %v has type: %T \n", interger, interger)
	fmt.Printf("Value %v has type: %T \n", f, f)
}

package main

import "fmt"

func main() {
	c1 := customErr{
		info: "This is my customized error",
	}
	foo(c1)
}

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	fmt.Sprintf("Here is the error: %v", ce.info)
	return ce.info
}

func foo(e error) {
	fmt.Println("Foo ran - ", e, "\n")
}

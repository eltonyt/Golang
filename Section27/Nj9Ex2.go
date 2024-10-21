package main

import "fmt"

type person struct {
	first string
	last  string
}

func (p *person) speak() {
	fmt.Printf("Speaking from %v %v. \n", p.first, p.last)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	el := person{
		first: "Elton",
		last:  "Li",
	}
	// saySomething(el) -> This won't work cuz speak() method has a pointer as receiver but el is a value here.
	saySomething(&el)
}

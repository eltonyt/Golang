package main

import "fmt"

type person struct {
	first string
	age   int
}

func (p person) speak() {
	fmt.Printf("Person %s is %d years old", p.first, p.age)
}

func main() {
	p := person{
		first: "Yoyo",
		age:   55,
	}
	p.speak()
}

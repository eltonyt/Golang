package main

import "fmt"

var (
	a, b, c *string
	d       *int
)

func init() {
	p := "Drop by drop, the bucket gets filled."
	q := "Persistently, patiently, you are bound to succeed."
	r := "The meaning of life is ..."
	n := 42
	a = &p
	b = &q
	c = &r
	d = &n
}

func main() {
	fmt.Printf("Address of a is: %v \n", a)
	fmt.Printf("Address of b is: %v \n", b)
	fmt.Printf("Address of c is: %v \n", c)
	fmt.Printf("Address of d is: %v \n", d)

	fmt.Printf("Type of a is: %T \n", a)
	fmt.Printf("Type of b is: %T \n", b)
	fmt.Printf("Type of c is: %T \n", c)
	fmt.Printf("Type of d is: %T \n", d)

	fmt.Printf("Value of a is: %v \n", *a)
	fmt.Printf("Value of b is: %v \n", *b)
	fmt.Printf("Value of c is: %v \n", *c)
	fmt.Printf("Value of d is: %v \n", *d)
}

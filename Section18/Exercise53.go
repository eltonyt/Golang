package main

import "fmt"

func main() {
	type person struct {
		firstName        string
		lastName         string
		favoriteICFlavor []string
	}

	p1 := person{firstName: "Elton", lastName: "Li", favoriteICFlavor: []string{"Chocolate"}}
	fmt.Printf("Person %s %s likes ", p1.firstName, p1.lastName)
	for _, v := range p1.favoriteICFlavor {
		fmt.Printf("%s ice cream flavor ", v)
	}

	p2 := person{firstName: "Yolo", lastName: "Li", favoriteICFlavor: []string{"Straberry", "Vanilla"}}
	fmt.Printf("Person %s %s likes ", p2.firstName, p2.lastName)
	for _, v := range p2.favoriteICFlavor {
		fmt.Printf("%s ice cream flavor ", v)
	}
}

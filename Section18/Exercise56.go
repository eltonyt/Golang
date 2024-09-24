package main

import "fmt"

func main() {
	a := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first:     "Elton",
		friends:   map[string]int{"test": 123, "real": 456},
		favDrinks: []string{"Coke", "Pepsi", "Coffee"},
	}

	fmt.Println(a)
}

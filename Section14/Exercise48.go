package main

import "fmt"

func main() {
	slice := [][]string{}
	slice = append(slice, []string{"James", "Bond", "Shaken, not stirred"})
	slice = append(slice, []string{"Miss", "MoneyPenny", "I'm 008", "123123"})
	fmt.Printf("Slice is %v", slice)
}

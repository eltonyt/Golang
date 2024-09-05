package main

import "fmt"

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	if x, ok := m["Q"]; ok {
		fmt.Printf("Value for key Q is %d.\n", x)
	} else {
		fmt.Printf("Value for key Q does not exist.")
	}
}

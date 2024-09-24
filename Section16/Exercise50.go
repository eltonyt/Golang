package main

import "fmt"

func main() {
	m := map[string][]string{
		"bind_james":       []string{"shaken, not stirred", "martinis", "fast cars"},
		"moneypenny_jenny": []string{"intelligence", "literature", "computer science"},
		"no_dr":            []string{"cats", "ice cream", "sunsets"},
	}

	m["fleming_ian"] = []string{"steaks", "cigars", "espionage"}

	for person, loveItems := range m {
		fmt.Printf("%v loves items ", person)
		for i, v := range loveItems {
			fmt.Printf("Index %d %v ;", i, v)
		}
		fmt.Print("\n")
	}
}

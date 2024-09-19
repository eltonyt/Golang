package main

import "fmt"

func main() {
	states := make([]string, 0, 50)
	states = append(states, "AL", "AK", "AZ", "AR", "CA", "CO", "CT", "DE", "FL", "GA", "HI", "ID", "IL", "IN", "IA", "KS", "KY", "LA", "ME", "MD", "MA", "MI", "MN", "MS", "MO", "MT", "NE", "NV", "NH", "NJ", "NM", "NY", "NC", "ND", "OH", "OK", "OR", "PA", "RI", "SC", "SD", "TN", "TX", "UT", "VT", "VA", "WA", "WV", "WI", "WY")
	fmt.Printf("Length of the slice: %d \n", len(states))
	fmt.Printf("Cap of the slice: %d \n", cap(states))
	for i := 0; i < len(states); i++ {
		fmt.Printf("Value at index %d is %v \n", i, states[i])
	}

}

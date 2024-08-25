package main

import "fmt"

// EXERCISE 4
type value int64

const (
	_        = iota // ignore first value by assigning to blank identifier
	KB value = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	// ZB
	// YB
)

func main() {
	fmt.Printf("%d \t %b\n", KB, KB)
	fmt.Printf("%d \t %b\n", MB, MB)
	fmt.Printf("%d \t %b\n", GB, GB)
	fmt.Printf("%d \t %b\n", TB, TB)
	fmt.Printf("%d \t %b\n", PB, PB)
	fmt.Printf("%d \t %b\n", EB, EB)
	// fmt.Printf("%d \t %b\n", ZB, ZB)
	// fmt.Printf("%d \t %b\n", YB, YB)
}

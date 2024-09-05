package main

import "fmt"

// Part A
// func main() {
// 	for i := range 100 {
// 		fmt.Printf("%d \n", i)
// 	}
// }

// Part B
func main() {
	for i := 0; i < 100; i++ {
		loop()
	}
}
func loop() {
	for i := range 100 {
		fmt.Printf("%d \n", i)
	}
}

package main

import "fmt"

func main() {
	var zero int
	zeroShort := 0
	multipleValue1, multipleValue2 := 1, 2
	var specifiedVar int = 222
	e, f, _ := 4, 5, 6
	fmt.Printf("Zero Value: %d \n", zero)
	fmt.Printf("Zero ValueShort Declarion: %d \n", zeroShort)
	fmt.Printf("Multiple Value 1: %d, Multiple Value 2: %d \n", multipleValue1, multipleValue2)
	fmt.Printf("Specified Value: %d \n", specifiedVar)
	fmt.Printf("E: %d \n", e)
	fmt.Printf("F: %d \n", f)
}

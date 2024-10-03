package main

import "fmt"

func main() {
	x := func1()
	y1, y2 := func2()
	fmt.Printf("X is %d\n", x)
	fmt.Printf("Y1 is %s\n, Y2 is %s", y1, y2)
}

func func1() int {
	return 1
}

func func2() (int, string) {
	return 1, "wth"
}

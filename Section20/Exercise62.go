package main

import (
	"fmt"
	"math"
)

type SQUARE struct {
	length int
	width  int
}

type CIRCLE struct {
	radius float64
}

func (s SQUARE) area() float64 {
	return float64(s.length * s.width)
}

func (c CIRCLE) area() float64 {
	return math.Pi * c.radius * c.radius
}

type SHAPE interface {
	area() float64
}

func INFO(s SHAPE) {
	fmt.Printf("The area of the shape is %v \n", s.area())
}

func main() {
	sq := SQUARE{
		length: 5,
		width:  6,
	}
	c := CIRCLE{
		radius: 2.5,
	}
	INFO(sq)
	INFO(c)
}

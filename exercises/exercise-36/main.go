package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 { //method for finding the area of square
	a := s.length * s.width
	return a
}

func (c circle) area() float64 { // method for finding the area of a circle
	b := math.Pi * c.radius * c.radius
	return b
}

type shape interface { // this is the interface whatever satisify this signature it align with this interface
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	sq := square{
		length: 3,
		width:  5,
	}
	ci := circle{
		radius: 7,
	}

	info(sq)
	info(ci)
}

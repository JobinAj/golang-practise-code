package main

import (
	"fmt"
	"math"
)

func main() {
	x := powernator(3)
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
}

func powernator(a float64) func() float64 {
	var c float64
	return func() float64 {
		c++
		return math.Pow(a, c) // this is the closure function because the outer funcion execution has finished and the variable is still being used and values are returned so this is a closure function.
	}
}

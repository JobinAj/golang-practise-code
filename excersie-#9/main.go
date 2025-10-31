package main

import (
	"fmt"
)

var z int = 20

const y int = 33

func main() {
	x := 40
	fmt.Printf("the value of x is %d and the type of x is %T\n", x, x)
	fmt.Printf("the value of y is %d and the type of y is %T\n", y, y)
	fmt.Printf("the value of z is %d and the type of z is %T\n", z, z)
}

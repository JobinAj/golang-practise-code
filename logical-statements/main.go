package main

import (
	"fmt"
)

func main() {
	x := 20
	y := 40

	if x < 42 && x < 50 {
		fmt.Printf("both should be less then the actual value to return the true boolean value\n")
	}

	if x > 10 || y < 50 {
		fmt.Printf("Any one condition should be satisified as we are using the OR operator\n")
	}

	if x != 30 && y != 41 {
		fmt.Printf("we are using NOT operator here\n")
	}
}

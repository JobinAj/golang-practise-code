package main

import (
	"fmt"
)

func main() {
	y := 7
	for i := 0; i < 20; i++ {
		fmt.Printf("the value of %v\n", i)
	}

	for y < 10 {
		fmt.Printf("second for loop the value of %v\n", y)
		y++
	}

	y = 7

	for {
		fmt.Printf("y is %v \t\t third for loop\n", y)
		if y > 20 {
			break
		}
		y++
	}

	for i := 0; i < 20; i++ {
		if i%3 != 0 {
			continue
		}
		fmt.Printf("counting even numbers %v\n", i)

	}
}

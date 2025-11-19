package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i <= 100; i++ {
		println(i)
	}
	for i := 0; i <= 100; i++ {

		x := rand.Intn(10)
		y := rand.Intn(10)
		/*if x < 4 && y < 4 {
		fmt.Printf("x and y are less than 4\n")
		} else if y > 6 && x > 6 {
			fmt.Printf("x and y is greater than 6\n")

			} else if x >= 4 && x <= 6 {
				fmt.Printf("x is greater than or equal to 4 and x is less than or equal to 6\n")
				} else {
					fmt.Printf("none of the previous case were met\n")
					}*/

		switch {
		case x < 4 && y < 4:
			fmt.Printf("x and y are less than 4\n")
		case y > 6 && x > 6:
			fmt.Printf("x and y is greater than 6\n")
		case x >= 4 && x <= 6:
			fmt.Printf("x is greater than or equal to 4 and x is less than or equal to 6\n")
		default:
			fmt.Printf("none of the previous case were met\n")
		}
	}
}

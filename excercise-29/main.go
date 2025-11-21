package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf("this outer loopis %v\t the inner loop is %v\n", i, j)
			// fmt.Printf("this is printed by the inner loop %v\n", j)

		}
		fmt.Println("")

	}
}

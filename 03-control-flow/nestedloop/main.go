package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("---")
		for j := 0; j < 5; j++ {
			fmt.Printf("the outer loop %v \t the inner loop is %v\n", i, j)
		}
	}
}

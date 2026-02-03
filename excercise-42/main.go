package main

import (
	"fmt"
)

func main() {
	/*
		xi := [5]int{10, 11, 12, 13, 14}

		for i, v := range xi {

			fmt.Printf("the index is %v\t the value is %v\n", i, v)
		}
	*/
	xi := [5]int{}

	for i := 0; i < 5; i++ {
		xi[i] = i

	}
	for i, v := range xi {
		fmt.Printf("the value is %v\t\n the index is %v\t\n the type is %T\n", i, v, v)
	}

}

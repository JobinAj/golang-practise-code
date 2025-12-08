package main

import "fmt"

func main() {
	xi := [5]int{10, 11, 12, 13, 14}

	for i, v := range xi {

		fmt.Printf("the index is %v\t the value is %v\n", i, v)
	}

}

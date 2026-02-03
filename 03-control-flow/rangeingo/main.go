package main

import (
	"fmt"
)

func main() {
	a := []int{47, 393, 9834, 3484}
	for i, v := range a {
		fmt.Printf("the index of  %v is \t %v\n", i, v)
	}

	b := map[string]int{
		"james": 42,
		"fred":  43,
	}
	for k, v := range b {
		fmt.Printf("the key is %v \t and the value of that particular key is %v\n", k, v)
	}
}

package main

import "fmt"

func main() {
	xi := []int{}

	for i := 42; i <= 51; i++ {
		xi = append(xi, i)
	}

	for _, v := range xi {
		fmt.Printf("the value is %v\tAnd the type is %T\n", v, v)
	}

}

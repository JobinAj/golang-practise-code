package main

import (
	"fmt"
	"math/rand"
)

func game(a, b int) {
	num1 := rand.Intn(a)
	num2 := rand.Intn(b)
	fmt.Printf("\nthe number is %d\n", num1)
	fmt.Printf("\nthe number is %d\n", num2)
	fmt.Printf("\nthe number is %d\n", num1)
	fmt.Printf("\nthe number is %d\n", num2)
	fmt.Printf("\nthe number is %d\n", num1)
}

package main

import (
	"fmt"
)

func hello(a, b string) string {
	sum := a + b
	fmt.Printf("the sum of %s and %s", a, b)
	return sum
}

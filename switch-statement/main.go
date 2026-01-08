package main

import (
	"fmt"
)

func main() {
	var x int
	println("enter the number:")
	fmt.Scanln(&x)

	switch x {
	case 43:
		fmt.Println("x is 43")
		fallthrough
	case 44:
		fmt.Println("printed next value because of fallthrough x is 44")
	case 45:
		fmt.Println("x is 45")
	case 40:
		fmt.Println("x is 40")
	default:
		fmt.Println("doesn't matches any cases")
	}
}

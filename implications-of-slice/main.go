package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6, 7}
	b := a

	fmt.Println(a)
	fmt.Println(b)
	println("-------------------")

	a[0] = 10

	fmt.Println(a)
	fmt.Println(b)
	println("-------------------")

}

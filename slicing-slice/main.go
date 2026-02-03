package main

import "fmt"

func main() {
	xi := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Printf("xi - %#v\n", xi[:])
	fmt.Printf("xi - %#v\n", xi[0:])
	fmt.Printf("xi - %#v\n", xi[5:])
	fmt.Printf("xi - %#v\n", xi[:6])

}

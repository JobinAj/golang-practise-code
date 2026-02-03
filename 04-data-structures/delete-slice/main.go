package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("%#v\n", xi)
	xi = append(xi[:3], xi[4:]...) // slice using index value i mean the value that you see inside the square brackets
	fmt.Printf("%#v\n", xi)
}

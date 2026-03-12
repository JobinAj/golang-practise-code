package main

import "fmt"

// pointer referance type
func DeltaInt(n *int) {
	*n = 43
}

// slice referance type
func DeltaSlice(n []int) {
	n[0] = 99
}

func main() {
	//passing value to the function intDelta
	a := 42
	fmt.Println(a)
	DeltaInt(&a)
	fmt.Println(a)

	//passing value to the underlying array
	xi := []int{1, 2, 3, 4}
	fmt.Println(xi)
	DeltaSlice(xi)
	fmt.Println(xi)

}

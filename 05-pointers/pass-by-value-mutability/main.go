package main

import (
	"fmt"
)

// pointer referance type
func DeltaInt(n *int) {
	*n = 43
}

// slice referance type
func DeltaSlice(n []int) {
	n[0] = 99
}

// map referance type
func DeltaMap(md map[string]int, nextyear string) {
	md[nextyear] = 23
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

	//passing value to the map referance type
	m := make(map[string]int)
	m["jobin"] = 22
	fmt.Println(m["jobin"])
	DeltaMap(m, "jobin")
	println(m["jobin"])
}

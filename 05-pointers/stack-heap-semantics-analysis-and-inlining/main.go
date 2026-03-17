package main

import "fmt"

// This function will be inlined by complier as it is a small function this function itself doesn't executes the compiler just creates a temp and calculates the values and prints the output.
func Add(n int) int {
	n += 1
	return n
}

func main() {
	fmt.Println(Add(1)) //here it escape to heap because not because of the value or the function just because of the argument 'println' so the 1 will go to heap
}

/* everything is done just to save the function overhead
and for cpu and memory effieiency brillient compiler
designed by brillients ken thompsen */

package main

import "fmt"

func main() {
	a := [2]int{10, 20}
	fmt.Println(a)

	b := [...]float64{3.14, 1.7826}
	fmt.Println(b)

	var c [2]int
	c[0] = 19
	c[1] = 30

	for i, v := range c {

		fmt.Printf("the index is %v type %T\t the value is %v type %T\t", i, v, i, v)
	}

}

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := 10
	if z := 20 * rand.Intn(x); z >= x {
		fmt.Printf("%v: z is greater than x: %v", z, x)
	} else {
		fmt.Printf("%v:z is lesser than x:%v", z, x)
	}

}

package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	m := map[float32]string{
		3.14:    "pi",
		2.71828: "eulers rule",
	}
	if v, ok := m[2.71828]; ok {
		fmt.Println(v)
	}

	fmt.Println("------------------")

	//statements statements idiom
	c := 0
	for i := 0; i <= 100; i++ {
		if x := rand.IntN(5); x == 3 {
			fmt.Printf("the count is %v \t the iteration is %v\t the value of x is %v \n", c, i, x)

		}
		c++
	}

}

package main

import (
	"fmt"
)

func main() {
	b := 2 % 30
	c := 2 / 30
	fmt.Printf("the reminder of 2 module 30 is %v\n and 2 divided by 30 is %v\n", b, c)

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Printf("counting the even numbers from one to hundred %v\n", i)
		}
	}
}

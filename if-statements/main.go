package main

import (
	"fmt"
)

var z = 20

func init() {
	fmt.Printf("static value is :%d\n", z)
}
func main() {
	if z < 21 {
		fmt.Printf("value of z is actually:%d\n", z)

	}
	if z < 21 {
		fmt.Printf("value of z is actually:%d\n", z)
	} else {
		fmt.Printf("the changed value of z is actually triggerd the else block:%d\n", z)
	}

	if z < 21 {
		fmt.Printf("value of z is actually:%d\n", z)
	} else if z == 21 {
		fmt.Printf("it is equal:%d\n", z)
	} else {
		fmt.Printf("the changed value of z is actually triggerd the else block:%d\n", z)
	}

}

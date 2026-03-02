package main

import "fmt"

func main() {
	f := outer() /*when this runs it hold a value of func outer which is
	               return func() int {
		                return 23
	                  } */
	fmt.Println(f()) //now when we call f() it will look up on the return 23 now it will print 23
}

func outer() func() int {
	return func() int {
		return 23
	}
}

package main

import "fmt"

func main() {
	f := incrementor()
	fmt.Println(f()) //1
	fmt.Println(f()) //2
	fmt.Println(f()) //3
	fmt.Println(f()) //4
	fmt.Println(f()) //5
	fmt.Println(f()) //6
	fmt.Println(f()) //7
}

func incrementor() func() int { //this is called closure function because the value x decalred in the incrementor stays alive and gets incremented when used in the inner function
	x := 0
	return func() int {
		x++
		return x
	}
}

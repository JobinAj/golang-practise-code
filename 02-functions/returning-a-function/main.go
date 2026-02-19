package main

import "fmt"

func main() {
	fmt.Println(foo()) // here function calls foo and foo returns 42 and prints 42 and the value is not stored anywhere.
	x := foo()
	fmt.Println(x) // but in here the function calls foo and foo returns 42 and stores it in x and we print x "In here we are storing the value for the purpose of reusing it"
	y := faint()
	fmt.Println(y()) // but in here the function calls foo and foo returns 42 and stores it in x and we print x "In here we are storing the value for the purpose of reusing it"

}

func foo() int {
	return 42
}

func faint() func() int {
	return func() int { // function returning function and it returns value
		return 45
	}
}

package main

import "fmt"

func main() {

	func(s string) {
		fmt.Println(s) // this is anonymous-function
	}("hey interfaces is really easy dude")

	golang := func(b string) { // this is function expression and expressions are anything like operator,varibles,function that produces single value.
		fmt.Println(b)
	}

	golang("hey golang is the cloud-native programming language so learn it")
}

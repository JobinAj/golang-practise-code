package main

import "fmt"

func main() {
	foo() // so this is a named function which implies to identifier
	func() {
		fmt.Println("so if we don't use HSTS then be ready to be getting hacked")
	}()

	func(s string) {
		fmt.Println(s)
	}("so give first preference to security")
}

func foo() {
	fmt.Println("Hey there use HSTS in your network protocols")
}

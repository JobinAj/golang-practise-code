package main

import "fmt"

//basic function

func main() {
	x := foo()
	fmt.Println(x)
	i, s := bar()
	fmt.Println(i, s)

}

func foo() int {
	return 42
}

func bar() (int, string) {
	return 43, "james bond"
}

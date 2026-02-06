package main

import "fmt"

func main() {
	a1, b1 := 12, 15
	a2, b2 := 193, 12
	x := doMath(a1, b1, Add)
	y := doMath(a2, b2, Substract)
	fmt.Printf("the Addition of %d & %d is :%d\n", a1, b1, x)
	fmt.Printf("the substraction of %d & %d is :%d", a2, b2, y)

}

func doMath(a, b int, f func(int, int) int) int {
	return f(a, b)
}

func Add(a, b int) int {
	return a + b
}

func Substract(a, b int) int {
	return a - b
}

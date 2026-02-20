package main

import "fmt"

func main() {
	x := DoMath(12, 15, Add)
	y := DoMath(36, 30, Substract)

	fmt.Printf("%T\n", Add)
	fmt.Printf("%T\n", Substract)
	fmt.Printf("%T\n", DoMath)

	fmt.Printf("The addition is:%v\nThe substraction is:%v", x, y)

}

func DoMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func Add(a int, b int) int {
	return a + b
}

func Substract(a int, b int) int {
	return a - b
}

package main

import "fmt"

func main() {
	fmt.Println(factorial(4))
	fmt.Println(factloop(4))

}

func factorial(n int) int {
	fmt.Printf("This is the value of n %d\n", n)
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
func factloop(n int) int {
	total := 1
	for n > 0 {
		total *= n
		n--
	}
	return total
}

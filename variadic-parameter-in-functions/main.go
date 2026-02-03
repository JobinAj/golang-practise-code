package main

import "fmt"

func main() {
	fmt.Println("sum of 1,2,3:", sum(1, 2, 3))
	fmt.Println("sum of 2,5:", sum(2, 5))
	fmt.Println("sum of nothing:", sum())

}

func sum(num ...int) int {
	total := 0
	for _, n := range num {
		total += n
	}
	return total
}

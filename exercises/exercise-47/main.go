package main

import "fmt"

func main() {
	fmt.Println(factorial(4))
	fmt.Println(factloop(4))

}

func factorial(n int) int {
	fmt.Println("this is n", n)
	if n == 1 {
		return 1
	}
	return n * factorial(n-1) //this is called reursion running the face same function again and again until it meets some condition.
}

// we can do the same thing with loop
func factloop(n int) int {
	total := 1
	for n > 0 {
		total *= n
		n--
	}
	return total
}

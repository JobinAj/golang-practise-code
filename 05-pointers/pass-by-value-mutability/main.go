package main

import "fmt"

//pointer referance type
func DeltaInt(n *int) {
	*n = 43
}

func main() {
	a := 42
	fmt.Println(a)
	DeltaInt(&a)
	fmt.Println(a)

}

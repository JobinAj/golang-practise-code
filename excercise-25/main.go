package main

import "math/rand"

func main() {
	for i := 0; i <= 42; i++ {
		println(i)
		x := rand.Intn(5)
		switch x {
		case 0:
			println("the value of x is: 0")
		case 1:
			println("the value of x is: 1")
		case 2:
			println("the value of x is: 2")
		case 3:
			println("the value of x is: 3")
		case 4:
			println("the value of x is: 4")
		default:
			println("the value of x never goes beyond 5 so default case is useless")

		}
	}

}

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(300)
	fmt.Printf("The random Value generated is:%v\n", x)

	if x <= 100 {
		fmt.Printf("the value is less than 100\nTHE VALUE:%v\n", x)
	} else if x > 100 && x <= 200 {
		fmt.Printf("the value between 100&200 is\nTHE VALUE:%v\n", x)
	} else if x > 200 && x <= 250 {
		fmt.Printf("the value between 200&250 is\nTHE VALUE:%v\n", x)
	} else {
		fmt.Printf("this was more than 250\nTHE VALUE IS:%v\n", x)
	}
	fmt.Println(rand.Intn(4))
	fmt.Println(rand.Intn(4))
	fmt.Println(rand.Intn(4))
	fmt.Println(rand.Intn(4))
	fmt.Println(rand.Intn(4))
	fmt.Println(rand.Intn(4))
	fmt.Println(rand.Intn(4))
	fmt.Println(rand.Intn(4))
	fmt.Println(rand.Intn(4))
	fmt.Println(rand.Intn(4))

}

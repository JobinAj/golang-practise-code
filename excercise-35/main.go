package main

import "fmt"

func main() {
	xs := []string{"Almond Biscotti Café", "Banana Pudding", "Balsamic Strawberry (GF)"}
	fmt.Printf("len(xs)=%d\n", len(xs))
	fmt.Printf("%T\n", xs)
	for _, v := range xs {
		fmt.Printf("%v\n", v)
	}
	fmt.Println("---------------------")

	println(xs[0])
	println(xs[1])
	println(xs[2])

	fmt.Println("---------------------")

	for i := 0; i < len(xs); i++ {
		fmt.Println(xs[i])

	}
}

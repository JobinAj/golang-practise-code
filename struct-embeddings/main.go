package main

import "fmt"

type sports_car struct {
	lamborgini string
}

type cool_car struct {
	sports_car
	bmw string
}

func main() {
	cars := cool_car{
		sports_car: sports_car{
			lamborgini: "urus",
		},
		bmw: "x1",
	}
	fmt.Println(cars)
}

package main

import (
	"fmt"

	lion "github.com/JobinAj/golang-modularizartion-learning"
)

func main() {
	s1 := lion.Cub()
	s2 := lion.Cubs()
	s3 := lion.Lion()
	s4 := lion.Lions()

	fmt.Printf("%s\n", s1)
	fmt.Printf("%s\n", s2)
	fmt.Printf("%s\n", s3)
	fmt.Printf("%s\n", s4)
}

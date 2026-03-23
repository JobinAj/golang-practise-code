package main

import "fmt"

var (
	a, b, c *string
	d       *int
)

func init() {
	p := "Drop by drop, the bucket gets filled."
	q := "Persistently, patiently, you are bound to succeed."
	r := "The meaning of life is ..."
	n := 42
	a = &p
	b = &q
	c = &r
	d = &n
}

func main() {
	fmt.Printf("the type of a is %v,%T\n", a, a)
	fmt.Printf("the type of b is %v,%T\n", b, b)
	fmt.Printf("the type of c is %v,%T\n", c, c)
	fmt.Printf("the type of d is %v,%T\n", d, d)

	//dereferance and print the value of the memory address that the four variable is pointing to

	fmt.Println(*a)
	fmt.Println(*b)
	fmt.Println(*c)
	fmt.Println(*d)
}

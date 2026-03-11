package main

import "fmt"

func main() {
	v := 20
	fmt.Println(v)
	x := &v // this for refering the address
	fmt.Println(x)
	*x = 50 //this is for derefering the value of that address
	fmt.Println(*x)
}

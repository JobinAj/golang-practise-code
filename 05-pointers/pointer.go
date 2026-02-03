package main

import "fmt"

func Pointers() {
	var i = "jobinaj"
	p := &i
	fmt.Printf("the address is %p", p)
}

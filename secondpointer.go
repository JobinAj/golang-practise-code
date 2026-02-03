package main

import "fmt"

func secondPointers() {
	var i = "jobinaj"
	p := &i
	fmt.Printf("the address is %p", p)
}

package main

import "fmt"

func bitshift() {
	fmt.Printf("\n%d %b", 1, 1)
	fmt.Printf("\n%d %b", 1<<1, 1>>1)
	fmt.Printf("\n%d %b", 1<<2, 1<<2)
	fmt.Printf("\n%d %b", 1<<3, 1<<3)
	fmt.Printf("\n%d %b", 1<<4, 1<<4)
	fmt.Printf("\n%d %b", 1<<5, 1<<5)
	fmt.Printf("\n%d %b", 1<<6, 1<<6)

}

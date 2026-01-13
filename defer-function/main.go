package main

import "fmt"

func main() {
	defer fmt.Printf("hello") // this will print only after the surronding function ends/
	fmt.Printf("world\t")
}

package main

import "fmt"

func main() {
	xj := []string{"James", "Bond", "Martini", "Chocolate"}
	xg := []string{"Jenny", "Moneypenny", "Guiness", "Chocolate"}
	fmt.Println(xj)
	fmt.Println(xg)
	xi := [][]string{xj, xg}
	fmt.Println(xi)
}

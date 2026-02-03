package main

import "fmt"

func main() {
	m := map[string]int{
		"jobin": 22,
		"alvin": 26,
	}
	for x, y := range m {
		fmt.Printf("name %v:age %v\n", x, y)
	}
}

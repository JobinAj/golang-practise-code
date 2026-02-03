package main

import "fmt"

func main() {
	m := map[string]int{
		"jobin": 22,
		"alvin": 26,
	}
	if v, ok := m["jobin"]; ok { //anything inside a if statement is false then the if statement will not run
		fmt.Printf("the value of lookup of jobin is %v", v)

	}
}

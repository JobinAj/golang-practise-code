package main

import "fmt"

func main() {
	records := [][]string{{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "I'm 008."},
	}

	for i, v := range records {
		fmt.Printf("the index is %v\t the value is %v\n", i, v)
		for a, b := range v {
			fmt.Printf("the index is %v\t the value is %v\n", a, b)
		}
	}
}

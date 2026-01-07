package main

import "fmt"

var tests = []struct {
	name  string
	input int8
	want  int8
}{
	{"double", 2, 4},
	{"double", 3, 6},
}

func main() {
	for _, t := range tests {
		fmt.Println(t.name, t.input, t.want)
	}
}

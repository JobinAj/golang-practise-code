package main

import "fmt"

type person struct {
	first string
	age   int
}

func (p person) speak() {
	fmt.Println("the firstname is", p.first, "my age is ", p.age)
}

func main() {
	p1 := person{
		first: "jobin",
		age:   23,
	}
	p1.speak()
}

package main

import "fmt"

type person struct {
	first string
	age   int
}

func (p person) speak() { //p is the reciver so it recive the value from the p1 when we call speak
	fmt.Println("the firstname is", p.first, "my age is ", p.age)
}

func main() {
	p1 := person{
		first: "jobin",
		age:   23,
	}
	p1.speak() // now the p1 value is passed to the p in the method and now p has the value that we created in the
}

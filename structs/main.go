package main

import "fmt"

type person struct {
	first_name string
	last_name  string
	age        int8
}

func main() {
	human := person{
		first_name: "jobin",
		last_name:  "aj",
		age:        23,
	}
	fmt.Printf("the firstname is %#v and the lastname is %#v and the age of the particular person is %v", human.first_name, human.last_name, human.age)
}

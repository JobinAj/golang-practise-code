package main

import "fmt"

//Define an interface

type speaker interface {
	speak() string
}

//Now define the concrete type the struct

type human struct {
	Name string
}

func (h human) speak() string {
	return "my name is " + h.Name
}

func makespeak(s speaker) {
	fmt.Println(s.speak())
}

func main() {
	h := human{Name: "jobin"}
	makespeak(h)

}

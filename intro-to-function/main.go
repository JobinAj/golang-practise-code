package main

import "fmt"

//main function
func main() {
	foo()
	bar("Jobin")
	age := personaldetails(24)
	fmt.Println("age", age)
	company, role := designation("sirpi", "devops engineer")
	fmt.Println("company", company, "role", role)
}

//function with no parm and no return value
func foo() {
	fmt.Println("foo")
}

//function with one parm and no return value
func bar(name string) {
	fmt.Println("my name is", name)
}

//function with one parm and one return value
func personaldetails(age int8) int8 {
	return age
}

//function with multiple parm and multiple return value
func designation(company string, role string) (string, string) {
	return company, role
}

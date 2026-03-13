package main

import "fmt"

// semantic values
func Add(n int) int {
	n += 1
	return n
}

// semantic pointer
func Deltasemantic(v *int) {
	*v += 1
}

func main() {
	//semantic values (pass by value)
	v := 1
	fmt.Println(v)      //here it will print 1
	fmt.Println(Add(v)) //here we are passing a copy of the v so v original value will be still 1 so here it will print 2 but next time if i print v it will 1.
	fmt.Println(v)
	fmt.Println("------------------------------------")
	//semantic pointer (pass by referance)
	a := 1
	fmt.Println(a)
	Deltasemantic(&a)
	fmt.Println(a)
	fmt.Println(a) //i printed again see the orginal value changed(magic😂)
}

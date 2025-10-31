// this is how we print hello world in go
package main //this is tells go that this is a executable code not just a file

import (
	"fmt"
	"hello/new"
) // we can also us println --> 'P' must be always in uppercase."Println is a builtin function which is used for debugging purposes but the best practises is using the Println through the fmt package"

func main() {
	fmt.Println("hello world")
	greetings()
	jobin()
	characterencode()
	new.Literals()
	Variables()
	hexadecimalandbinary()
	Variable_scope()
	Type_Conversion()
	Notype()
	game(100, 200)
	hello("two", "four")
	a, b := swap("hello", "world")
	fmt.Printf("\n%s,%s", a, b)
	d, c := noreturn(42)
	fmt.Printf("\n%d %d \n", d, c)
	h, g := shortdeclaration(false, true)
	fmt.Printf("%t\t%t\t", h, g)
	novaluevar()
	varwithvalue()
	basictypes()
	infernaceandconstants()
	bitshift()
	increment()
}
func jobin() {
	sum := new.Add(5, 2)
	sub := new.Substract(200, 30)
	carspeed, myname, bullet := datatypes()
	fmt.Println("what the world")
	fmt.Printf("The sum is %v\n", sum)
	fmt.Printf("The sub is %d\n", sub)
	fmt.Printf("my name is %s,and i drive car at %d,and the bullet speed was %f.\n", myname, carspeed, bullet)

}

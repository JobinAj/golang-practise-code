// this is how we print hello world in go
package main //this is tells go that this is a executable code not just a file

import (
	"fmt"
) // we can also us println --> 'P' must be always in uppercase."Println is a builtin function which is used for debugging purposes but the best practises is using the Println through the fmt package"

func main() {
	fmt.Println("hello world")
	greetings()
	jobin()
	characterencode()
	Literals()
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

func greetings()                              {}
func characterencode()                        {}
func Literals()                               {}
func Variables()                              {}
func hexadecimalandbinary()                   {}
func Variable_scope()                         {}
func Type_Conversion()                        {}
func Notype()                                 {}
func game(a, b int)                           {}
func hello(a, b string)                       {}
func swap(a, b string) (string, string)       { return b, a }
func noreturn(a int) (int, int)               { return a, a }
func shortdeclaration(a, b bool) (bool, bool) { return a, b }
func novaluevar()                             {}
func varwithvalue()                           {}
func basictypes()                             {}
func infernaceandconstants()                  {}
func bitshift()                               {}
func increment()                              {}
func datatypes() (int, string, float64)       { return 0, "", 0.0 }

func jobin() {
	sum := Add(5, 2)
	sub := Substract(200, 30)
	carspeed, myname, bullet := datatypes()
	fmt.Println("what the world")
	fmt.Printf("The sum is %v\n", sum)
	fmt.Printf("The sub is %d\n", sub)
	fmt.Printf("my name is %s,and i drive car at %d,and the bullet speed was %f.\n", myname, carspeed, bullet)

}

func Add(a, b int) int       { return a + b }
func Substract(a, b int) int { return a - b }

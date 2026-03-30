package main

import "fmt"

type person struct {
	first string
}

func main() {
	p := person{
		first: "jobin",
	}
	fmt.Println(p)
	p = changename(p, "alvin") //value semantics
	fmt.Println(p)
	changenameptr(&p, "God is a awesome God") //pointer semantics
	fmt.Println(p)

}

func changename(p person, s string) person {
	p.first = s
	return p
}

func changenameptr(p *person, s string) {
	p.first = s
}

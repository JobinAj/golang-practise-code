package main

import "fmt"

type dog struct { // this is a struct
	firstname string
}

func (d dog) walk() {
	fmt.Println("My name is", d.firstname, "and i'm walking.")
}

func (d *dog) run() {
	d.firstname = "rover"
	fmt.Println("My name is", d.firstname, "and i'm running")
}

func main() {
	d1 := dog{"henry"}
	d1.walk()
	d1.run() // go automatically converts this to (&d10).run()

	d2 := &dog{"padget"}
	d2.walk()
	d2.run()

}

/* Methods works just simliar to the function methods converts everything in the same way how function works.
in the background but we use methods so that we can attach behaviour to the type and also we can call the function
in a more meaningfull way*/

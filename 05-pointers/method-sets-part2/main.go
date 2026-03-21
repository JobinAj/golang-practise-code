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

type youngin interface {
	walk()
	run()
}

func youngrun(y youngin) {
	y.run()
}

func main() {
	d1 := dog{"henry"}
	d1.walk()
	d1.run() // go automatically converts this to (&d10).run()
	//youngrun(d1) //but go automatically converts in the case of method calls only if it is interface go does'nt converts

	d2 := &dog{"padget"}
	d2.walk()
	d2.run()
	youngrun(d2)

}

/* Methods works just simliar to the function methods converts everything in the same way how function works.
in the background but we use methods so that we can attach behaviour to the type and also we can call the function
in a more meaningfull way*/

//when we add interface to the method sets.

/* while calling the methods sets the go automatically changes the d1.run() to (&d1).run().
but in the interface it's not the same the go doesn't automatically changes it to the "*,or &"
so when we do youngrun(d1) so the d1 goes to the youngin and takes the place the of y but
then youngin ask hey dude stop right there do you integral the interface youngin first meaning do you have both the run() and walk()
then d1 check yeah i just now checked i can see dog in the walk but for run it is *dog sorry so that function never executes*/

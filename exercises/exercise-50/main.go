package main

import "fmt"

type dog struct { // this is a struct
	firstname string
}

func (d dog) walk() {
	fmt.Println("My name is", d.firstname, "and i'm walking.")
}

func (d dog) run() {
	d.firstname = "rover"
	fmt.Println("My name is", d.firstname, "and i'm running")
}

type youngin interface {
	walk()
	run()
}

func youngrun(y youngin) {
	y.run()
	y.walk()
}

func main() {
	d1 := dog{"henry"}
	youngrun(d1) //but go automatically converts in the case of method calls only if it is interface go does'nt converts

	d2 := dog{"padget"}
	youngrun(d2)
	d2 = d2.changename("Rover")
	youngrun(d2)

}

//just for fun changing the name

func (d dog) changename(s string) dog { // here when we do d2.changename we pass d2 to the d  now a new varible of same type dog has the copy of the value stored in d2 now we did d.firstname=s so we just passed rover so now the value changed to rover and now we have to return d so that when we call d2 the value should be changed.
	d.firstname = s
	return d
}

/* Regarding this excersie i have to implement both the methods to work with the youngin interface i can just change the youngrun(&d1)
but the value semantics is the first preferance and also the recommendation is only using pointer semantics if the value is larger more thatn 64bytes.
so now i have to make it work without the pointer semantics so just i have to change the mehods type to value semantics to the both the methods,
and i have to remove the address refereance in the d2 main.*/

/* Methods works just simliar to the function methods converts everything in the same way how function works.
in the background but we use methods so that we can attach behaviour to the type and also we can call the function
in a more meaningfull way*/

//when we add interface to the method sets.

/* while calling the methods sets the go automatically changes the d1.run() to (&d1).run().
but in the interface it's not the same the go doesn't automatically changes it to the "*,or &"
so when we do youngrun(d1) so the d1 goes to the youngin and takes the place the of y but
then youngin ask hey dude stop right there do you integral the interface youngin first meaning do you have both the run() and walk()
then d1 check yeah i just now checked i can see dog in the walk but for run it is *dog sorry so that function never executes*/

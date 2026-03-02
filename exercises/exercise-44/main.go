package main

import "fmt"

func main() {
	fmt.Println(square(3))
}

func printsquare(f func(int) int, a int) string { //here f is a varibale and its type is func so any func that satify this signature then it can be used here,
	x := f(a) /*when it runs f=square because square has this type signature this is the callback function,
	  and also this is where the actual callback happens because  printsquare did created the function we just passed square to the to the printsquare function and it recived it as f so now f is calling back square. so that is callback function*/
	return fmt.Sprintf("the sqaured is %v the value is %v", a, x) //using sprint because we are returning string
}

func square(n int) int {
	return n * n //this is where the actual logic is.
}

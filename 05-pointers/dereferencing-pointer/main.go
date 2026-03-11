package main

import "fmt"

func main() {
	j := 101 //new variable declared
	fmt.Println(j)
	o := &j //so now "o" has the address of j
	fmt.Println(o)
	*o = 100         // now we are replacing the value 101 with 100 so now
	fmt.Println(*o)  //printing the value of pointer to "o"
	fmt.Println(j)   //now priting the value of j because "o" was the downstream address referance of j so now idomatically the j value is changed successfully
	fmt.Println(&*o) //checked by printing the address it has the same address as "j"
	fmt.Println(&j)  //checked by printing the address it has the same address as "o"
	/*note: meaning the variable names are same the address is same so value can be changed*/
}

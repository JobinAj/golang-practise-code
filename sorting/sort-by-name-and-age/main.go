package main

import (
	"fmt"
)

type Person struct{
	Age int 
	Name string
}

func (p Person) String() string{
        return fmt.Sprintf("%s:,%d",p.Age,p.Name)
}





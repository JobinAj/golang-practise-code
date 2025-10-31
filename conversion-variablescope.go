package main

import (
	"fmt"
)

var Starlink int64 = 1000000000

func Variable_scope() {
	var asustuf string = "gaming laptop"
	fmt.Printf("global variables is:%d\n", Starlink)
	fmt.Printf("local variables is:%s\n", asustuf)

}

func Type_Conversion() float64 {
	convert := float64(Starlink)
	fmt.Printf("converted %d to %f", Starlink, convert)
	return convert

}

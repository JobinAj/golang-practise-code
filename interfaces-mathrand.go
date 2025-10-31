package main

import (
	"fmt"
)

func Notype() {
	var anytype interface{}
	anytype = "jobinaj"
	fmt.Printf("\nmy name is:%s", anytype)
}

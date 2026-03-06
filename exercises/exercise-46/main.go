package main

import (
	"fmt"
	"time"
)

func main() {
	timefunc(dowork)
}

func dowork() {
	for i := 0; i < 2000; i++ {
		fmt.Println(i)
	}
}

func timefunc(f func()) {
	start := time.Now()
	f() //so this is a callback function and also a wrapper function because it adds extra behaviour to the function and how it is a callback function meaning it passes dowork as argument to the function and uses it later.
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

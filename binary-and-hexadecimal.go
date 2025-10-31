package main

import "fmt"

func hexadecimalandbinary() {
	var a int = 8327
	var b string = "jobin"
	var c string = "🤣"
	var d float32 = 872.923

	fmt.Printf("decimal:%d\t,%s\t,%s\t,%f\n", a, b, c, d)
	fmt.Printf("hexadecimal:%X\t,%X\t,%X\t,%X\n", a, b, c, d)
	fmt.Printf("binary:%b\t,%s\t,%s\t,%b\n", a, b, c, d)
}

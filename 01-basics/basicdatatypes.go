package main

import "fmt"

var (
	strings         = "magic"
	float32s        = 94.00
	bools           = true
	uints           = 2392   //only positive numbers
	complexs        = 3 + 4i // four is real and i is imaginery
	names    []rune = []rune{'j', 'o', 'b', 'i', 'n'}
)

func basictypes() {
	fmt.Printf("\nstrings is %s:", strings)
	fmt.Printf("\nfloat is %f:", float32s)
	fmt.Printf("\nboolean is %t:", bools)
	fmt.Printf("\nnon-negative is %d:", uints)
	fmt.Printf("\nrunes is %s:", string(names))
	fmt.Printf("\ncomplex is %v:", complexs)
}

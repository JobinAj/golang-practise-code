package main

import (
	"fmt"
)

func increment() {
	const (
		_ = iota
		a
		b
		c
		d
		e
		jack
		son
	)
	fmt.Printf("\n\n%d %d %d %d %d %d %d", a, b, c, d, e, jack, son)
}

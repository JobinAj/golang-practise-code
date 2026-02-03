package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	d1 := time.Duration(rand.Int63n(250))
	d2 := time.Duration(rand.Int63n(250))

	go func() {
		time.Sleep(d1 * time.Millisecond)
		ch1 <- "i'm first"
	}()

	go func() {
		time.Sleep(d2 * time.Millisecond)
		ch2 <- "i'm first"
	}()

	select {
	case v1 := <-ch1:
		fmt.Println("value from the channel 1", v1)
	case v2 := <-ch2:
		fmt.Println("value from the channel 2", v2)
	}
}

package main

import (
	"fmt"
	"log"
	"strconv"
)

type book struct {
	title string
}

func (b book) String() string {
	return fmt.Sprintln("The Title of the book is ", b.title)
}

type count int

func (c count) String() string {
	return fmt.Sprint("this is the number", strconv.Itoa(int(c)))
}

func loginfo(s fmt.Stringer) {
	log.Println("LOG OF LINE 138", s.String())
}

func main() {
	b := book{
		title: "The Gospel of john",
	}

	var n count = 25

	loginfo(b)
	loginfo(n)
}

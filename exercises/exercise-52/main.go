package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Firstname  string
	Secondname string
	Age        int
}

func main() {
	p1 := person{
		Firstname:  "jobin",
		Secondname: "aj",
		Age:        23,
	}
	p2 := person{
		Firstname:  "alvin",
		Secondname: "aj",
		Age:        26,
	}
	people := []person{p1, p2}

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println("error", err)
	}

	//os.Stdout.Write(bs) //
	fmt.Println(string(bs))

}

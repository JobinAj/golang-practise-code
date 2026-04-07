package main

import (
	"encoding/json"
	"fmt"
)

type person []struct {
	Firstname  string `json:"Firstname"`
	Secondname string `json:"Secondname"`
	Age        int    `json:"Age"`
}

func main() {
	sb := []byte(`[{"Firstname":"jobin","Secondname":"aj","Age":23},{"Firstname":"alvin","Secondname":"aj","Age":26}]`) //this is just a string data not directly useable in the go so we create a struct above and saying to go that except json like this and map to this shape.

	//create a variable to store the result

	var p person

	err := json.Unmarshal(sb, &p)
	if err != nil {
		fmt.Println("the error is:", err)
		return
	}
	fmt.Println(p)
}

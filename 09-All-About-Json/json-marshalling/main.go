package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// created struct which is accessiable outside of this module
type ColourGroup struct {
	ID     int
	Name   string
	Colors []string
}

// main function starts
func main() {
	group := ColourGroup{
		ID:     1,
		Name:   "Redis",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	/*here what we are doing is b,err:= type inferance
	so what ever json.Marshal returns b will be of that type so the marshall signature is []byte*/
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b) //we using stdout here because if we use println here it will print raw byte because b has slices of byte []byte so we are using stdout.write to get the orginal json values

	var Newgroup ColourGroup

	err = json.Unmarshal(b, &Newgroup)
	if err != nil {
		println("error", err)
	}
	fmt.Println("\nthe unmarshalled data of same b marshaled data is", Newgroup)

}

package main

import "fmt"

func main() {
	location := Paradise("Heaven")
	fmt.Println(location)

}

func Paradise(location string) string {
	return fmt.Sprintf("The location where i think paradise is %s", location)
}

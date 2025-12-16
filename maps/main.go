package main

import "fmt"

func main() {
	people := map[string]int{
		"jobin aj": 22,
		"alvin aj": 26,
	}
	fmt.Println("the people age is", people["jobin aj"])
	fmt.Println("the length of the map is", len(people))
	fmt.Println(people)

	timemachinevalues := make(map[string]int)

	timemachinevalues["i don't know the values actually that's why i'm using make😅"] = 10
	timemachinevalues["i don't know the values actually that's why i'm using make i don't know second thing also😅😅"] = 20

	fmt.Println(timemachinevalues["i don't know the values actually that's why i'm using make😅"])
	fmt.Println(timemachinevalues["i don't know the values actually that's why i'm using make i don't know second thing also😅😅"])

	people["albert"] = 59
	people["jyothi"] = 51

	fmt.Println(people)

	for k, v := range people {
		fmt.Println(k, v)
	}
	for k, _ := range people {
		fmt.Println(k)
	}
	//this is comma ok idiom
	// v, ok := people["jobinaj"]
	// if ok {
	// 	fmt.Println(v)
	// } else {
	// 	fmt.Println("the value doesn't exist")
	// }
	if v, ok := people["alvin aj"]; !ok {
		fmt.Println("the value did'nt exist")
	} else {
		fmt.Println(v)
	}
}

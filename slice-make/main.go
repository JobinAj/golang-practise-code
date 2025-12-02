package main

import "fmt"

func main() {
	/* slice append using slice literals */
	xi := []int{0, 1, 2, 3, 5, 6, 7}
	xi = append(xi, 8, 9, 10)
	fmt.Printf("%v\n", xi)

	println("-----------------------")

	/*slice append using slice make */

	mi := make([]int, 0, 20)
	fmt.Println(mi)
	fmt.Println(len(mi))
	fmt.Println(cap(mi))
	mi = append(mi, 10, 20, 30, 40, 60, 70, 80, 90, 100, 110, 120, 130, 140, 150, 160, 170, 190, 200, 210, 220, 230)
	fmt.Println(mi)
	fmt.Println(len(mi))
	fmt.Println(cap(mi))

}

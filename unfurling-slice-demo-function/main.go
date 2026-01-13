package main

func main() {
	num := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	xi := unfurl(num...)
	print(xi)
}

func unfurl(ii ...int) int {
	if len(ii) == 0 {
		return 0
	}
	return ii[0] + unfurl(ii[1:]...)
}

package main

import (
	"fmt"
	"sort"
)

func main(){
xi:=[]int{4,7,3,42,99,18,56,12}
xs:=[]string{"james","Q","Moneypenny","Dr. No"}
fmt.Println(xi)
fmt.Println(xs)
sort.Ints(xi)
fmt.Printf("used the standard library to print the sorted slice of integer:%v\n",xi)
sort.Strings(xs)
fmt.Printf("used the standard library to print the sorted slice of string:%s",xs)
}

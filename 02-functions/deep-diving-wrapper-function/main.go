package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	xb, err := readfile("poem.txt")
	if err != nil {
		log.Fatalf("error in the main in the readfile:%s", err)
	}
	fmt.Println(xb)
	fmt.Println(string(xb))

}

func readfile(poem string) ([]byte, error) {
	xb, err := os.ReadFile(poem)
	if err != nil {
		return nil, fmt.Errorf("there was a error in the readfile:%s", err)
	}
	return xb, nil
}

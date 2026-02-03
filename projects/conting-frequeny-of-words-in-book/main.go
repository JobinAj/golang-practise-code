package main

import (
	"bufio" //
	"fmt"   //formatted printing the statementsfmt
	"os"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("book.txt")
	if err != nil {
		fmt.Println("error opening file", err)
		return
	}
	defer file.Close()

	//map to store wordfrequency
	wordcount := make(map[string]int)

	//store the bytes in buffer and read the token's line by line

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		//convert the text to lower case
		line = strings.ToLower(line)

		//remove puncuations and only include letters and space

		line = strings.Map(func(r rune) rune {
			if unicode.IsLetter(r) || unicode.IsSpace(r) {
				return r
			}
			return -1

		}, line)
		//split lines into words
		words := strings.Fields(line)

		//count the words
		for _, word := range words {
			wordcount[word]++
		}
	}
	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Print word frequencies
	for word, count := range wordcount {
		fmt.Printf("%s : %d\n", word, count)
	}
}

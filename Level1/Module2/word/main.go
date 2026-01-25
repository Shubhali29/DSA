package main

import (
	"fmt"
	"strings"
)

func main() {

	var word string

	fmt.Scan(&word)
	lowerCount := 0
	upperCount := 0

	for _, value := range word {

		if value >= 'a' && value <= 'z' {
			lowerCount = lowerCount + 1
		}

		if value >= 'A' && value <= 'Z' {
			upperCount = upperCount + 1
		}
	}

	if lowerCount >= upperCount {
		fmt.Println(strings.ToLower(word))
	} else if lowerCount < upperCount {
		fmt.Println(strings.ToUpper(word))
	}
}

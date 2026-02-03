package main

import (
	"fmt"
	"strings"
)

/*
Given two strings s and t, determine if they are isomorphic.

Two strings s and t are isomorphic if the characters in s can be replaced to get t.

All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character, but a character may map to itself.

*/

func main() {

	var input string
	fmt.Scan(&input)

	var finalString []string

	for _, value := range input {

		lowerValue := strings.ToLower(string(value))
		if lowerValue == "a" || lowerValue == "e" || lowerValue == "i" || lowerValue == "o" || lowerValue == "u" || lowerValue == "y" {
			continue
		}
		finalString = append(finalString, ".")
		finalString = append(finalString, lowerValue)
	}

	fmt.Println(strings.Join(finalString, ""))
}

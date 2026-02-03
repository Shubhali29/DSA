package main

import (
	"fmt"
	"strings"
)

/*
Petya started to attend programming lessons. On the first lesson his task was to write a simple program. The program was supposed to do the following: in the given string, consisting if uppercase and lowercase Latin letters, it:

deletes all the vowels,
inserts a character "." before each consonant,
replaces all uppercase consonants with corresponding lowercase ones.
Vowels are letters "A", "O", "Y", "E", "U", "I", and the rest are consonants. The program's input is exactly one string, it should return the output as a single string, resulting after the program's processing the initial string.

Help Petya cope with this easy task.
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

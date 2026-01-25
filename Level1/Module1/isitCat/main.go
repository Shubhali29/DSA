package main

import (
	"fmt"
	"strings"
)

const (
	catSound = "meow"
)

// For example, strings "meow", "mmmEeOWww", "MeOooOw" describe a meowing, but strings "Mweo", "MeO", "moew", "MmEW", "meowmeow" do not.

func checkSound(length int, soundString string) string {

	if length < 4 {
		return "NO"
	}

	if strings.ToLower(string(soundString[0])) != "m" {
		return "NO"
	}

	lowerCaseString := strings.ToLower(soundString)
	var updatedString string
	var foundM, foundE, foundO, foundW bool
	// remove repeated characters
	for _, value := range lowerCaseString {

		if value != 'm' && value != 'e' && value != 'o' && value != 'w' {
			return "NO"
		}

		if value == 'm' {
			if foundM {
				continue
			}
			updatedString = updatedString + string(value)
			foundM = true
		}

		if value == 'e' {
			if foundE {
				continue
			}
			updatedString = updatedString + string(value)
			foundE = true
			foundM = false
		}

		if value == 'o' {
			if foundO {
				continue
			}
			updatedString = updatedString + string(value)
			foundO = true
			foundE = false
			foundM = false
		}

		if value == 'w' {
			if foundW {
				continue
			}
			updatedString = updatedString + string(value)
			foundW = true
			foundO = false
			foundE = false
			foundM = false
		}
	}

	if updatedString == catSound {
		return "YES"
	}

	return "NO"
}
func main() {

	var testCases int
	var result []string

	fmt.Scanln(&testCases)

	for i := 1; i <= testCases; i++ {

		var length int
		var soundString string
		fmt.Scanln(&length)
		fmt.Scanln(&soundString)

		answer := checkSound(length, soundString)

		result = append(result, answer)

	}

	for _, value := range result {
		fmt.Println(value)
	}
}

/*
7
4
meOw
14
mMmeoOoWWWwwwW
3
mew
7
MmeEeUw
4
MEOW
6
MmyaVW
5
meowA
*/

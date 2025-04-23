package main

import (
	"fmt"
	"strings"
)

func main() {

	var FirstNamePerson1, LastNamePerson1, FirstNamePerson2, LastNamePerson2 string

	fmt.Println("Enter First Name and Last Name of Person 1:")
	fmt.Scanln(&FirstNamePerson1, &LastNamePerson1)
	fmt.Println("Enter First Name and Last Name of Person 2:")
	fmt.Scanln(&FirstNamePerson2, &LastNamePerson2)

	if strings.EqualFold(LastNamePerson1, LastNamePerson2) {
		fmt.Println("Brothers")
	} else {
		fmt.Println("Not Brothers")
	}
}

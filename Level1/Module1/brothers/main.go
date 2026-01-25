package main

import (
	"fmt"
	"strings"
)

func main() {

	var fnPerson1, snPerson1, fnPerson2, snPerson2 string

	fmt.Scanln(&fnPerson1, &snPerson1)
	fmt.Scanln(&fnPerson2, &snPerson2)

	if strings.ToLower(snPerson1) == strings.ToLower(snPerson2) {
		fmt.Println("ARE Brothers")
	} else {
		fmt.Println("NOT")
	}

}

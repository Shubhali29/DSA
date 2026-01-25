package main

import (
	"fmt"
)

func main() {

	var numberOfPerson int

	fmt.Scan(&numberOfPerson)

	opinion := make([]int, numberOfPerson)

	for i := 0; i < numberOfPerson; i++ {
		fmt.Scan(&opinion[i])
	}

	for i := 0; i < numberOfPerson; i++ {

		if opinion[i] == 1 {
			fmt.Println("HARD")
			return
		}
	}

	fmt.Println("EASY")

}

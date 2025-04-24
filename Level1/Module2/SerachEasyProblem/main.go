package main

import (
	"fmt"
)

func main() {

	var totalPerson int
	fmt.Println("Enter the total number of persons: ")
	fmt.Scan(&totalPerson)

	opinion := make([]int, totalPerson)

	for i := 1; i <= totalPerson; i++ {
		fmt.Println("Enter the opinion (1 for yes, 0 for no): ")
		fmt.Scan(&opinion[i-1])
	}

	for i := 0; i < len(opinion); i++ {

		if opinion[i] == 1 {
			fmt.Println("HARD")
			return
		}
	}

	fmt.Println("EASY")
}

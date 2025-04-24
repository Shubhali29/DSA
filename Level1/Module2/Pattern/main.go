package main

import (
	"fmt"
)

func main() {

	var rows int
	fmt.Println("Enter the number of rows: ")
	fmt.Scan(&rows)

	if rows <= 2 || rows > 20 {
		fmt.Println("Invalid input")
		return
	}

	// Print the first row of stars
	for i := 1; i <= rows; i++ {
		fmt.Print("*")
	}

	fmt.Println()
	k := 1

	// Print the middle rows of stars
	for i := 2; i < rows; i++ {
		for j := 1; j < (rows - k); j++ {
			fmt.Print(" ")
		}
		fmt.Print("*")
		k++
		fmt.Println()

	}

	// Print the last row of stars
	for i := 1; i <= rows; i++ {
		fmt.Print("*")
	}

}

// Problem statement
// You are given a number 'n'.

// Return number of digits in ‘n’.

// Example:
// Input: 'n' = 123

// Output: 3

// Explanation:
// The 3 digits in ‘123’ are 1, 2 and 3.

// Detailed explanation ( Input/output format, Notes, Images )

package main

import (
	"fmt"
)

//  100/10,

func iterative(num int) int {

	noOfDigits := 0

	for num > 0 {

		num = num / 10
		noOfDigits++
	}

	return noOfDigits
}

func recursion(num int) int {

	if num <= 0 {
		return 0
	}
	return recursion(num/10) + 1
}

func main() {

	var num int
	fmt.Println("Enter number")
	fmt.Scanln(&num)

	noOfDigits := iterative(num)

	fmt.Printf("Number of digits (Iterative): %d", noOfDigits)

	noOfDigits = recursion(num)

	fmt.Printf("\n Number of digits (recursion): %d", noOfDigits)
}

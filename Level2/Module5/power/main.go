// Problem statement
// Write a program to find x to the power n (i.e. x^n). Take x and n from the user. You need to return the answer.

// Do this recursively.

// Detailed explanation ( Input/output format, Notes, Images )
// Input format :
// Two integers x and n (separated by space)
// Output Format :
// x^n (i.e. x raise to the power n)

package main

import (
	"fmt"
)

// 2^2 = 2*2
func iterative(x, n int) int {

	result := 1

	for n > 0 {
		result *= x
		n--
	}

	return result
}

func recursive(x, n int) int {

	if n == 0 {
		return 1
	}
	return x * recursive(x, (n-1))
}

func main() {

	var x, n int
	fmt.Println("Enter number")
	fmt.Scanln(&x)
	fmt.Println("Enter power")
	fmt.Scanln(&n)

	result := iterative(x, n)

	fmt.Printf("Result %d", result)

	result = recursive(x, n)

	fmt.Printf("\n Result (Recursive) %d", result)
}

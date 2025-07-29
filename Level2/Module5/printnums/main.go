// Problem statement
// Given the number 'n', write a code to print numbers from 1 to n in increasing order recursively.

// Detailed explanation ( Input/output format, Notes, Images )
// Input Format :
// Integer n
// Output Format :
// Numbers from 1 to n (separated by space)

package main

import (
	"fmt"
)

func iterative(n int) {
	count := 1

	for n > 0 {
		fmt.Println(count)
		count++
		n--
	}
}

func recursion(n int, initial int) {

	if n == 0 {
		return
	}

	fmt.Println(initial)
	recursion(n-1, initial+1)
}
func main() {

	var n int
	fmt.Println("Enter value of n")
	fmt.Scanln(&n)

	fmt.Println("Iterative")
	iterative(n)
	fmt.Println("Recusive")
	recursion(n, 1)

}

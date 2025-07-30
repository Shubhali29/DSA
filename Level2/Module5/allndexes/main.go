// Problem statement
// Given an array of length N and an integer x, you need to find all the indexes where x is present in the input array. Save all the indexes in the output array (in increasing order).

// Do this recursively. Indexing in the array starts from 0.

// Hint:
// Try making a helper function with the required arguments and call the helper function from the allIndexes function.
// Detailed explanation ( Input/output format, Notes, Images )
// Input Format :
// Line 1 : An Integer N i.e. size of array
// Line 2 : N integers which are elements of the array, separated by spaces
// Line 3 : Integer x
// Output Format :
// Return all the indexes in the output array (in increasing order).
// Constraints :
// 1 <= N <= 10^3
package main

import (
	"fmt"
)

func iterative(list []int, x int) []int {

	var indexes []int

	for index, value := range list {

		if x == value {
			indexes = append(indexes, index)
		}
	}
	return indexes
}

func recursive(list []int, x int, index int, result *[]int) {

	if index >= len(list) {
		return
	}

	if list[index] == x {
		*result = append(*result, index)
	}

	recursive(list, x, index+1, result)
}

func main() {

	list := []int{10, 9, 5, 1, 9, 1}
	fmt.Println("List is %v", list)
	var x int
	fmt.Println("Enter number to find last index in list")
	fmt.Scanln(&x)

	allIndex := iterative(list, x)
	fmt.Println("Iterative")
	fmt.Printf("\n All Index of %d is present %d", x, allIndex)
	var result []int
	recursive(list, x, 0, &result)

	fmt.Println("\n Recursive")
	fmt.Printf("\n All Index of %d is present %d", x, result)

}

// Problem statement
// Given an array of length N and an integer x, you need to find and return the first index of integer x present in the array. Return -1 if it is not present in the array.

// First index means, the index of first occurrence of x in the input array.

// Do this recursively. Indexing in the array starts from 0.

// Detailed explanation ( Input/output format, Notes, Images )
// Input Format :
// Line 1 : An Integer N i.e. size of array
// Line 2 : N integers which are elements of the array, separated by spaces
// Line 3 : Integer x
// Output Format :
// first index or -1
// Constraints :
// 1 <= N <= 10^3

package main

import (
	"fmt"
)

func iterative(list []int, x int) int {
	firstIndex := -1

	for index, value := range list {

		if value == x {
			firstIndex = index
			break
		}
	}
	return firstIndex
}

func recursive(list []int, x, index int) int {

	if index >= len(list) {
		return -1
	}

	if x == list[index] {
		return index
	}

	return recursive(list, x, index+1)
}

func main() {

	list := []int{10, 9, 5, 1, 9, 1}
	fmt.Println("List is %v", list)
	var x int
	fmt.Println("Enter number to find first index in list")
	fmt.Scanln(&x)

	firstIndex := iterative(list, x)
	fmt.Println("Iterative")
	fmt.Printf("\n First Index of %d is present %d", x, firstIndex)

	firstIndex = recursive(list, x, 0)

	fmt.Println("\n Recursive")
	fmt.Printf("\n First Index of %d is present %d", x, firstIndex)

}

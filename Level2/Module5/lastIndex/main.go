// Problem statement
// Given an array of length N and an integer x, you need to find and return the last index of integer x present in the array. Return -1 if it is not present in the array.

// Last index means - if x is present multiple times in the array, return the index at which x comes last in the array.

// You should start traversing your array from 0, not from (N - 1).

// Do this recursively. Indexing in the array starts from 0.

// Detailed explanation ( Input/output format, Notes, Images )
// Input Format :
// Line 1 : An Integer N i.e. size of array
// Line 2 : N integers which are elements of the array, separated by spaces
// Line 3 : Integer x
// Output Format :
// last index or -1
// Constraints :
// 1 <= N <= 10^3

package main

import (
	"fmt"
)

func iterative(list []int, x int) int {

	for index := len(list) - 1; index >= 0; index-- {
		if list[index] == x {
			return index
		}
	}

	return -1
}

func recursive(list []int, x, index int) int {

	if index < 0 {
		return -1
	}

	if list[index] == x {
		return index
	}

	return recursive(list, x, index-1)
}

func recursive2(list []int, x, index int) int {

	if index >= len(list) {
		return -1
	}

	result := recursive(list, x, index+1)

	if result != -1 {
		return result
	} else if list[index] == x {
		return index
	}

	return -1
}

func main() {

	list := []int{10, 9, 5, 1, 9, 1}
	fmt.Println("List is %v", list)
	var x int
	fmt.Println("Enter number to find last index in list")
	fmt.Scanln(&x)

	firstIndex := iterative(list, x)
	fmt.Println("Iterative")
	fmt.Printf("\n Last Index of %d is present %d", x, firstIndex)

	firstIndex = recursive(list, x, len(list)-1)

	fmt.Println("\n Recursive")
	fmt.Printf("\n Last Index of %d is present %d", x, firstIndex)

	firstIndex = recursive2(list, x, 0)

	fmt.Println("\n Recursive")
	fmt.Printf("\n Last Index of %d is present %d", x, firstIndex)

}

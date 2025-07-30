// Problem statement
// Given an array of length N and an integer x, you need to find if x is present in the array or not. Return true or false.

// Do this recursively.

// Detailed explanation ( Input/output format, Notes, Images )
// Input Format :
// Line 1 : An Integer N i.e. size of array
// Line 2 : N integers which are elements of the array, separated by spaces
// Line 3 : Integer x
// Output Format :
// 'true' or 'false'
// Constraints :
// 1 <= N <= 10^3/

package main

import (
	"fmt"
)

func iterative(list []int, x int) bool {

	for _, value := range list {
		if value == x {
			return true
		}
	}
	return false
}

func recursive(list []int, x int, index int) bool {

	if index >= len(list) {
		return false
	}

	if x == list[index] {
		return true
	}

	return recursive(list, x, index+1)
}

func recursive2(list []int, x int, index int) bool {

	if index >= len(list) {
		return false
	}

	isPresent := recursive(list, x, index+1)

	if isPresent {
		return isPresent
	} else if x == list[index] {
		return true
	}

	return false
}

func main() {

	list := []int{10, 3, 5, 1, 9}

	var x int
	fmt.Println("Enter number to check in list")
	fmt.Scanln(&x)

	isPresent := iterative(list, x)
	fmt.Println("Iterative")
	if isPresent {
		fmt.Printf("\n %d is present", x)
	} else {
		fmt.Printf("\n %d is not present", x)
	}

	isPresent = recursive(list, x, 0)

	fmt.Println("\n Recursive")
	if isPresent {
		fmt.Printf("\n %d is present", x)
	} else {
		fmt.Printf("\n %d is not present", x)
	}

	isPresent = recursive2(list, x, 0)

	fmt.Println("\n Recursive2")
	if isPresent {
		fmt.Printf("\n %d is present", x)
	} else {
		fmt.Printf("\n %d is not present", x)
	}

}

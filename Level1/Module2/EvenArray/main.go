package main

import (
	"fmt"
)

func checkArrayEven(length int, array []int) int {

	// Check how many even nums and odd nums.
	// Check incorrect places

	incorrectPlaces := 0
	oddIndices, evenIndices := 0, 0
	oddNums, evenNums := 0, 0

	for i := 0; i < length; i++ {

		oddIndex, evenIndex, oddNum, evenNum := false, false, false, false

		if i%2 == 0 {
			evenIndex = true
			evenIndices += 1
		} else {
			oddIndex = true
			oddIndices += 1
		}

		if (array[i] % 2) == 0 {
			evenNum = true
			evenNums += 1
		} else {
			oddNum = true
			oddNums += 1
		}

		if !((evenIndex && evenNum) || (oddIndex && oddNum)) {
			incorrectPlaces += 1
		}
	}

	if incorrectPlaces == 0 {
		return 0
	}

	if (oddIndices != oddNums) || (evenIndices != evenNums) {
		return -1
	}

	return incorrectPlaces / 2
}

func main() {

	var testCases int

	fmt.Scan(&testCases)
	var result []int
	for i := 1; i <= testCases; i++ {

		var length int
		fmt.Scan(&length)
		array := make([]int, length)
		for i := 0; i < length; i++ {
			fmt.Scan(&array[i])
		}
		movesCount := checkArrayEven(length, array)

		result = append(result, movesCount)
	}

	for i := 0; i < testCases; i++ {
		fmt.Println(result[i])
	}
}

// 0 1 2 3 4 5
// 1 2 3 4 5 6

// 1 3 5 2 4 6
// W R W W R W

// 0 1 2 3
// 2 2 3 2

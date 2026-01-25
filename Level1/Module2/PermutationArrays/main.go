package main

import (
	"fmt"
)

/*
4
4 2 3 7
2 3 4 9
*/

func main() {

	var n int

	fmt.Scan(&n)

	arrayA := make([]int, n)
	arrayB := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arrayA[i])
	}

	for i := 0; i < n; i++ {
		fmt.Scan(&arrayB[i])
	}

	arrayAMap := make(map[int]int)
	arrayBMap := make(map[int]int)

	for i := 0; i < n; i++ {

		if _, exists := arrayAMap[arrayA[i]]; exists {
			arrayAMap[arrayA[i]] = arrayAMap[arrayA[i]] + 1
		} else {
			arrayAMap[arrayA[i]] = 1
		}

		if _, exists := arrayBMap[arrayB[i]]; exists {
			arrayBMap[arrayB[i]] = arrayBMap[arrayB[i]] + 1
		} else {
			arrayBMap[arrayB[i]] = 1
		}
	}

	for i := 0; i < n; i++ {
		if _, exists := arrayBMap[arrayA[i]]; exists {
			if arrayBMap[arrayA[i]] != arrayAMap[arrayA[i]] {
				fmt.Println("no")
				return
			}
		} else {
			fmt.Println("no")
			return
		}
	}

	fmt.Println("yes")
}

package main

import (
	"fmt"
)

func main() {

	var size int
	fmt.Println("Enter the size of the array: ")
	fmt.Scan(&size)

	arrayA := make([]int, size)
	arrayB := make([]int, size)

	for i := 0; i < size; i++ {
		fmt.Println("Enter the element of array A: ")
		fmt.Scan(&arrayA[i])
		fmt.Println("Enter the element of array B: ")
		fmt.Scan(&arrayB[i])
	}

	count := 0

	for i := 0; i < size; i++ {

		for j := i; j < size; j++ {
			if arrayA[i] == arrayB[j] {
				count++
				break
			}
		}
	}

	if count == size {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}

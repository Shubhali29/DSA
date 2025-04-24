package main

import (
	"fmt"
)

func main() {

	var totalNumber int
	var evenCount, oddCount, negativeCount, PositiveCount int

	fmt.Println("Enter the total number of elements: ")
	fmt.Scan(&totalNumber)

	numbers := make([]int, totalNumber)

	for i := 1; i <= totalNumber; i++ {
		fmt.Println("Enter the number: ")
		fmt.Scan(&numbers[i-1])
	}

	for i := 0; i < totalNumber; i++ {
		if numbers[i] < 0 {
			negativeCount++
		} else if numbers[i] > 0 {
			PositiveCount++
		}
		if numbers[i]%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}

	fmt.Println("Even : ", evenCount)
	fmt.Println("Odd : ", oddCount)
	fmt.Println("Positive : ", PositiveCount)
	fmt.Println("Negative : ", negativeCount)

}

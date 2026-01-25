package main

import (
	"fmt"
)

func main() {

	var length int
	var even, odd, positive, negative int

	fmt.Scan(&length)
	inputs := make([]int, length)

	for i := 0; i < length; i++ {
		fmt.Scan(&inputs[i])

	}

	for _, value := range inputs {
		if value < 0 {
			negative = negative + 1
		}

		if value > 0 {
			positive = positive + 1
		}

		if value%2 == 0 || value == 0 {
			even = even + 1
		} else {
			odd = odd + 1
		}
	}

	fmt.Println("Even: ", even)
	fmt.Println("Odd: ", odd)
	fmt.Println("Positive: ", positive)
	fmt.Println("Negative: ", negative)
}

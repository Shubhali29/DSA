package main

import (
	"fmt"
)

func main() {

	var number int
	fmt.Println("Enter the number: ")
	fmt.Scan(&number)

	fibancci := make([]int, number)
	fibancci[0] = 0
	fibancci[1] = 1

	for i := 2; i < number; i++ {
		fibancci[i] = fibancci[i-1] + fibancci[i-2]
	}

	fmt.Printf("Fibonacci number of %v: %v", number, fibancci[number-1])
}

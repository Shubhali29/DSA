package main

import (
	"fmt"
)

func factorial(number int) int {
	if number <= 0 {
		return 1
	}
	return number * factorial((number - 1))
}

func main() {
	var number int
	fmt.Println("Enter the number")
	fmt.Scanln(&number)

	fact := factorial(number)

	fmt.Printf("Factorial of %v is %v", number, fact)
}

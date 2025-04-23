package main

import (
	"fmt"
)

func main() {

	var num1, num2 int
	var sum int

	fmt.Println("Enter two numbers : ")
	fmt.Scanln(&num1, &num2)
	lastDigitNum1 := num1 % 10
	lastDigitNum2 := num2 % 10

	sum = lastDigitNum1 + lastDigitNum2
	fmt.Println("Sum of last digits is : ", sum)
}

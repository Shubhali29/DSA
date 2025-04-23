package main

import (
	"fmt"
)

func main() {

	var num1, num2 int
	var operator string
	fmt.Println("Enter an expression :")
	fmt.Scan(&num1, &operator, &num2)

	switch operator {
	case ">":
		if num1 > num2 {
			fmt.Println("Right")
		} else {
			fmt.Println("Wrong")
		}
	case "<":

		if num1 < num2 {
			fmt.Println("Right")
		} else {
			fmt.Println("Wrong")
		}
	case "=":
		if num1 == num2 {
			fmt.Println("Right")

		} else {
			fmt.Println("Wrong")
		}
	default:

		fmt.Println("Invalid operator")
	}
}

package main

import (
	"fmt"
)

func main() {

	var num1, num2 int
	// take input from user
	fmt.Printf("Enter two numbers : ")
	fmt.Scanln(&num1, &num2)

	if num1 >= num2 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}

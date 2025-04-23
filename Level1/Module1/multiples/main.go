package main

import (
	"fmt"
)

func main() {

	var num1, num2 int
	fmt.Println("Enter two numbers : ")
	fmt.Scanln(&num1, &num2)

	if num1%num2 == 0 || num2%num1 == 0 {
		fmt.Println("Multiples")
	} else {
		fmt.Println("No Multiples")
	}

}

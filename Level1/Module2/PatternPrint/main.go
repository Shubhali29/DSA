package main

import (
	"fmt"
)

func main() {

	var number int
	fmt.Println("Enter the number: ")
	fmt.Scan(&number)

	for i := 1; i <= number; i++ {

		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Print(" ")
	}

}

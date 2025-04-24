package main

import (
	"fmt"
)

func main() {

	var number int
	fmt.Println("Enter the number: ")
	fmt.Scan(&number)

	for i := 1; i <= number; i++ {
		if number%i == 0 {
			fmt.Println(i)
		}
	}
}

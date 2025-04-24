package main

import (
	"fmt"
)

func main() {

	var num int

	fmt.Println("Enter the number: ")
	fmt.Scan(&num)

	if num <= 1 {
		fmt.Println("NO")
		return
	}

	count := 0

	for i := 1; i <= num; i++ {

		if num%i == 0 {
			count++
		}

	}

	if count == 2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}

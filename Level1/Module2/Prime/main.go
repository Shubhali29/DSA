package main

import (
	"fmt"
)

func main() {

	var number int

	fmt.Scan(&number)

	if number <= 1 {
		fmt.Println("NO")
		return
	}

	factor := 1
	for i := 2; i <= number; i++ {

		if number%i == 0 {
			factor = factor + 1
		}

		if factor > 2 {
			break
		}
	}

	if factor == 2 {
		fmt.Println("YES")
		return
	}
	fmt.Println("NO")
}

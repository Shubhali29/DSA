package main

import (
	"fmt"
)

func main() {

	var a, b int

	fmt.Scan(&a)
	fmt.Scan(&b)

	lastDigitOfa := a % 10
	lastDigitOfb := b % 10

	sum := lastDigitOfa + lastDigitOfb

	fmt.Println(sum)

}

// 100 % 10 = 0
// 100 /10 = 10

// 123 % 10 = 12

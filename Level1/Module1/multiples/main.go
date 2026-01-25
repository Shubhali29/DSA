package main

import (
	"fmt"
)

func main() {

	var a, b int

	fmt.Scan(&a)
	fmt.Scan(&b)

	if a%b == 0 || b%a == 0 {
		fmt.Print("Multiples")
	} else {
		fmt.Print("No Multiples")
	}
}

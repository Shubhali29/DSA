package main

import (
	"fmt"
)

func main() {

	var a, b int

	fmt.Scan(&a)
	fmt.Scan(&b)

	if a >= b {
		fmt.Print("Yes")
	} else {
		fmt.Print("No")
	}
}

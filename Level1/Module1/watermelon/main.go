package main

import (
	"fmt"
)

func main() {

	var weight int
	fmt.Scanln(&weight)

	if weight%2 != 0 || weight == 2 {
		fmt.Println("NO")
	} else {

		fmt.Println("YES")
	}
}

package main

import (
	"fmt"
)

func main() {

	var a, b, k int
	fmt.Println("Enter three numbers : ")
	fmt.Scanln(&a, &b, &k)

	if a%k == 0 && b%k == 0 {
		fmt.Println("Both")
	} else if a%k == 0 {
		fmt.Println("Memo")
	} else if b%k == 0 {
		fmt.Println("Momo")
	} else {
		fmt.Println("None")
	}
}

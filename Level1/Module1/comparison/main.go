package main

import (
	"fmt"
)

func main() {

	var num1, num2 int
	var sign string

	fmt.Scan(&num1, &sign, &num2)

	switch {
	case sign == "<":
		if num1 < num2 {
			fmt.Println("Right")
		} else {
			fmt.Println("Wrong")
		}
	case sign == ">":
		if num1 > num2 {
			fmt.Println("Right")
		} else {
			fmt.Println("Wrong")
		}
	case sign == "=":
		if num1 == num2 {
			fmt.Println("Right")
		} else {
			fmt.Println("Wrong")
		}
	}

}

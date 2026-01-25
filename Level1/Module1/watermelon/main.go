package main

import (
	"fmt"
)

func main() {

	var watermelonWeight int

	fmt.Scan(&watermelonWeight)

	if watermelonWeight != 2 && watermelonWeight%2 == 0 {
		fmt.Print("Yes")
	} else {
		fmt.Print("No")
	}
}

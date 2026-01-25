package main

import (
	"fmt"
)

func main() {

	var player1, player2, dice int

	fmt.Scan(&player1, &player2, &dice)

	if (player1%dice == 0) && (player2%dice == 0) {
		fmt.Println("Both")
	} else if player1%dice == 0 {
		fmt.Println("Memo")
	} else if player2%dice == 0 {
		fmt.Println("Momo")
	} else {
		fmt.Println("No One")
	}

}

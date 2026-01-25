package main

import (
	"fmt"
	"strings"
)

func main() {

	var number, time int
	fmt.Scan(&number, &time)

	var queue string
	fmt.Scan(&queue)

	mutableQueue := make([]string, number)

	for i := 0; i < len(queue); i++ {
		mutableQueue[i] = string(queue[i])
	}

	for i := 1; i <= time; i++ {

		for j := 1; j < len(queue); j++ {

			if queue[j-1] == 'B' && queue[j] == 'G' {
				mutableQueue[j-1], mutableQueue[j] = string(queue[j]), string(queue[j-1])
			}
		}
		queue = strings.Join(mutableQueue, "")
	}

	fmt.Println(queue)
}

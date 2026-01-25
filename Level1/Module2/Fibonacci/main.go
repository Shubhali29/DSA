package main

import (
	"fmt"
)

func main() {

	var num int
	fmt.Scan(&num)

	if num == 1 {
		fib1 := 0
		fmt.Println(fib1)
		return
	}

	if num == 2 {
		fib2 := 1
		fmt.Println(fib2)
		return
	}

	fibSeries := make([]int, num)
	fibSeries[0] = 0
	fibSeries[1] = 1

	for i := 2; i < num; i++ {
		fibSeries[i] = fibSeries[i-1] + fibSeries[i-2]
	}

	fmt.Println(fibSeries[num-1])
}

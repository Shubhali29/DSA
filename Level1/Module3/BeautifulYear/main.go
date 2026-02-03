package main

import "fmt"

/*
It seems like the year of 2013 came only yesterday. Do you know a curious fact? The year of 2013 is the first year after the old 1987 with only distinct digits.

Now you are suggested to solve the following problem: given a year number, find the minimum year number which is strictly larger than the given one and has only distinct digits.

Input
The single line contains integer y (1000 ≤ y ≤ 9000) — the year number.

Output
Print a single integer — the minimum year number that is strictly larger than y and all it's digits are distinct. It is guaranteed that the answer exists.


*/

func checkYear(year int) bool {

	distinct := make(map[int]int)

	for year != 0 {

		remainder := year % 10
		year = year / 10

		if _, exists := distinct[remainder]; exists {
			return false
		} else {
			distinct[remainder] = 1
		}
	}

	if len(distinct) == 4 {
		return true
	}

	return false
}

func main() {

	var year int
	fmt.Scan(&year)

	if year < 1000 || year > 9000 {
		return
	}

	for {
		year = year + 1

		if checkYear(year) {
			fmt.Println(year)
			return
		}
	}
}

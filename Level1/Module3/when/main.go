package main

import (
	"fmt"
	"strconv"
)

/*

AtCoder Beginner Contest usually starts at 21:00 JST and lasts for
100 minutes.

You are given an integer
K between
0 and
100 (inclusive). Print the time
K minutes after 21:00 in the HH:MM format, where HH denotes the hour on the
24-hour clock and MM denotes the minute. If the hour or the minute has just one digit, append a
0 to the beginning to represent it as a
2-digit integer.

Constraints
K is an integer between
0 and
100 (inclusive).
*/

func countDigit(num int) int {

	count := 0

	if num == 0 {
		return 1
	}

	for num != 0 {
		count = count + 1
		num = num / 10
	}

	return count
}

func main() {

	var k int

	fmt.Scan(&k)

	currentTime := 21

	minutes := k % 60
	hour := k - minutes

	finalHour := currentTime + (hour / 60)

	hourCount := countDigit(finalHour)

	hourString := ""
	minuteString := ""

	if hourCount < 2 {
		hourString = "0" + strconv.Itoa(finalHour)
	} else {
		hourString = strconv.Itoa(finalHour)
	}

	minCount := countDigit(minutes)

	if minCount < 2 {
		minuteString = "0" + strconv.Itoa(minutes)
	} else {
		minuteString = strconv.Itoa(minutes)
	}

	fmt.Println(hourString + ":" + minuteString)
}

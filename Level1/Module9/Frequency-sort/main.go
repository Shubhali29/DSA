package main

import (
	"fmt"
	"sort"
	"strings"
)

type frequencySort struct {
	character string
	frequency int
}

func main() {
	var input string
	fmt.Scan(&input)
	frequencyMap := make(map[string]int)
	var frequencyList []frequencySort
	var sortedList []string

	for _, value := range input {
		value1 := string(value)
		if _, exists := frequencyMap[value1]; exists {
			frequencyMap[value1] = frequencyMap[value1] + 1
		} else {
			frequencyMap[value1] = 1
		}

	}

	for key, value := range frequencyMap {
		object := frequencySort{character: key, frequency: value}
		frequencyList = append(frequencyList, object)
	}

	// Custom sort
	sort.Slice(frequencyList, func(i, j int) bool {
		return frequencyList[i].frequency > frequencyList[j].frequency // Use '>' for decending
	})

	for _, value := range frequencyList {

		for j := 0; j < value.frequency; j++ {
			sortedList = append(sortedList, value.character)
		}
	}

	fmt.Println(strings.Join(sortedList, ""))
}

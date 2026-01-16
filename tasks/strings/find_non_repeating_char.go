// map[rune]int, strings.ToLower()


package main

import (
	"strings"
)

func FindNonRepeatingChar(s string) (rune) {

	// Converting the lowered string to runes
	lowered := strings.ToLower(s)
	runes := []rune(lowered)

	// Instantiating tempMap map[rune]int to keep track of substring count
	tempMap := make(map[rune]int, len(runes))

	for _, val := range runes {
		tempMap[val]++
	}

	for _, val := range lowered {
		if tempMap[val] == 1 {
			return val
		}
	}

	return rune(0)

}

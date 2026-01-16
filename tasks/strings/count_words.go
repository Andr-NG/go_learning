// strings.Fields(), make(map[string]int)


package main

import (
	"strings"
)

func CountWords(s string) map[string]int {

	splitString := strings.Fields(s)
    // wordCount := map[string]int{}
    wordCount := make(map[string]int)

    for _, val := range splitString {
        wordCount[string(val)]++
    }
    return wordCount
}
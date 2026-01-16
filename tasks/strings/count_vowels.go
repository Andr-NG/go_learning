package main

import "strings"

func CountVowels(s string) int {

	vowels := "aeiuoy"
	sLowered := strings.ToLower(s)

	var vowelCount int

	for _, val := range sLowered {
		if strings.ContainsRune(vowels, val) {
			vowelCount++
		}
	}
	return vowelCount
}
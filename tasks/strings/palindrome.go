// two-pointers approach, strings.ToLower()

package main

import (
	"strings"
)

func VerifyPalindrome(s string) bool {

	// Handling cases where the string has either 1 or 0 chars.
	sRunes := []rune(strings.ToLower(s))
	sLength := len(sRunes)
	if sLength < 2 {
		return true
	}

	// Using the two-pointer approach to compare runes.
	i, j := 0, len(sRunes)-1
	for i < j {
		if sRunes[i] != sRunes[j] {
			return false
		}
		i++
		j--
	}

	return true
}

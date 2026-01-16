// unicode.IsLetter(), unicode.IsDigit()


package main

import (
	"unicode"
)

func CountCharType(s string) map[string]int {

    resultingMap := map[string]int{}


	for _, val := range s {
        if unicode.IsLetter(val) {
            resultingMap["letters"]++
        } else if unicode.IsDigit(val) {
            resultingMap["digits"]++
        } else {
            resultingMap["special character"]++
        }

    }
    return resultingMap
}


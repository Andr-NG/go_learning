// strings.Repeat(s, count)

package main

import "strings"

func HideLastWord(s string) string {

	slicedString := strings.Split(s, " ")
    lastIndex := len(slicedString)-1
    toBeHidden := slicedString[lastIndex]

    slicedString[lastIndex] = strings.Repeat("*", len(toBeHidden))

    return strings.Join(slicedString, " ")
}
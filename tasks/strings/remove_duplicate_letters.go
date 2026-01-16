// make(map[rune]bool, strings.Builder, strings.Builder.WriteRune()



package main

import "strings"


func RemoveDuplicates(s string) string {

    runes := []rune(strings.ToLower(s))

    var stringBuilder strings.Builder

    charMap := make(map[rune]bool)

    for _, val := range runes {
        // Checking the key pressence
        _, ok := charMap[val]
        if !ok {
            // if not present, placing the key, assigning true and adding the char to the string
            charMap[val] = true
            stringBuilder.WriteRune(val)
        }

    }
    return stringBuilder.String()

}
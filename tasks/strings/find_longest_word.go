package main

import "strings"


func FindLongestWord(s string) string {

    slice := strings.Fields(s)
    if len(slice) == 1 {
        return s
    }
    maxLength := 0
    lengthMap := make(map[int]string)


    for _, val := range slice {
        wordLength := len([]rune(val))
        if wordLength > maxLength {
            lengthMap[wordLength] = val
            maxLength = wordLength
        }
    }
    return lengthMap[maxLength]
}



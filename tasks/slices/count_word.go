package main

import "strings"

func CountWord(s string) map[string]int {

	slicedS := strings.Split(s, " ")
    resultMap := make(map[string]int)

    for _, val := range slicedS {
        resultMap[val]++
    }

    return resultMap
}
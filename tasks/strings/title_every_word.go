package main

import (
	"strings"
)

func TitleEveryWord(s string) string {

    splitS := strings.Split(s, " ")
    
    var tempSlice []string

	for _, val := range splitS{
		val = strings.ToUpper(val)
        tempSlice = append(tempSlice, val)
	}
    return strings.Join(tempSlice, " ")
}

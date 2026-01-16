package replace_spaces

import (
	"strings"
)


func ReplaceSpaces() string {
	myString := "go is fun"
	var newString strings.Builder

	for _, val := range myString {
		if val == ' ' {
			newString.WriteString("-")

		} else{
			newString.WriteRune(val)
		}
	}
	return newString.String()
}
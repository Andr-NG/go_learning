package main



func ReverseString(s string) (string) {

	// Converting the string to runes to handle Unicode
	runes := []rune(s)

	// Defining the last index of the string
	lastIndex := len(runes) - 1
	i, j := 0, lastIndex
	for i < j {
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}

	return string(runes)

}
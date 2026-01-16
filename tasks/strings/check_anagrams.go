// make(map[rune]int), runeCount[val]--, runeCount[val]++ 

package main

func CheckAnagram(s1, s2 string) bool {

    // Defining slices of runes to cover all Unicode
	s1Runes := []rune(s1)
	s2Runes := []rune(s2)

    // Handling the case where the length is different returning false at once
	if len(s1Runes) != len(s2Runes) {
		return false
	}

    // Initialising a mapping to count runes
    runeCount := make(map[rune]int)

    for _, val := range s1Runes {
        runeCount[val]++
    }

    // Ranging over the 2 string to 
    for _, val := range s2Runes {
        runeCount[val]--
        if runeCount[val] < 0 {
            return false
        }
    }

    return true


}

package main

import "fmt"

func main() {

	// Initialising a mapping
	capitalMap := make(map[string]string)
	capitalMap["Spain"] = "Madrid"
	capitalMap["USA"] = "Washington D.C."
	capitalMap["Italy"] = "Rome"
	capitalMap["Germany"] = "Berlin"

	fmt.Println(capitalMap)

	// Printing capitals as keys
	for country := range capitalMap {
		fmt.Printf("Capital of %s is %s\n", country, capitalMap[country])
	}

	// Accessing non-existent key
	capital, ok := capitalMap["Mexico"]
	fmt.Println(capital, ok)

	if !ok {
		fmt.Println("No such a country")
	} else {
		fmt.Println("Capital is", capital)
	}

	// Deleting a key
	delete(capitalMap, "Germany")
	fmt.Println(capitalMap)
}
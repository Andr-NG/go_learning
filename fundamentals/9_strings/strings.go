package main

import (
	"fmt"
	"strings"
)


func main() {
	greeting := "Hello, world!"
	fmt.Printf("My programme says: \"%s\"\n", greeting)

	for i := 0; i < len(greeting); i++ {
		// fmt.Printf("Char %c = hex %x\n", greeting[i], greeting[i])
		// characters are represented in UTF-8 (and ASCII) by their decimal value.
		fmt.Println("Letter", string(greeting[i]), "at position", i)
	}

	fruits := []string{"apples", "bananas", "oranges"}
	fmt.Println(strings.Join(fruits, ", "))
}


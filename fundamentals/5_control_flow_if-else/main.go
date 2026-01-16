package main

import "fmt"

func main() {

	// age := 18

	// if age >= 18 {
	// 	fmt.Println("You are allowed to buy alcohol!")
	// } else {
	// 	fmt.Println("No alcohol for you")
	// }

	score := 86

	if score >= 90 {
		fmt.Println("Execelent")
	} else if score >= 80 {
		if score <= 85 {
			fmt.Println("Good")
		} else {
			fmt.Println("Great")
		}
	}

}
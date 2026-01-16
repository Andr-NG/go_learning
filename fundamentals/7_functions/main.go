package main

import "fmt"

func sayCapital(s string) {
	fmt.Printf("%s is a capital.\n", s)
}

func sayCountry(s string) {
	fmt.Printf("%s is a country.\n", s)
}

func isGreater(n1, n2 int) bool {
	return n1 > n2
}

func swapNames(surname, name string) (string, string) {
	return name, surname
}

func main() {
	sayCapital("Madrid")
	sayCountry("Spain")
	a, b := 100, 200
	fmt.Println(isGreater(a, b))
	name, surname := swapNames("Doe", "John")
	fmt.Printf("The correct name order is %s %s\n", name, surname)
}

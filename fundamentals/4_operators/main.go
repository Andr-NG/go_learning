// An operator is a symbol that performs specific mathematical manipulations.

// Arithmetic Operators
// Relational Operators
// Logical Operators
// Bitwise Operators
// Assignment Operators
// & Miscellaneous Operators

package main

import "fmt"

func main() {
	// Arithmetic Operators
	A := 10
	B := 20
	C := 11
	D := 3
	fmt.Println("A + B =", A+B) // Addition 30
	fmt.Println("A - B =", A-B) // Subtraction -10
	fmt.Println("A * B =", A*B) // Multiplication 200
	fmt.Println("A / B =", A/B) // Division 5.0
	fmt.Println("C % D =", C%D) // Modulus 2
	A++
	fmt.Println("A++ =", A) //Increment  11
	A--
	fmt.Println("A-- =", A) //Decrement  10
	// ********************************

	// Relational operators - Comparison Operators

	// Equal to
	fmt.Println("A == B:", A == B) // Equals - false

	//Not equal to
	fmt.Println("A != B:", A != B) // Equals - true

	// Greater than
	fmt.Println("A > B:", A > B) // Equals - false

	// Less than
	fmt.Println("A < B:", A < B) // Equals - true

	// Greater than or equal to
	fmt.Println("A >= B:", A >= B) // Equals - false

	// Less than or equal to
	fmt.Println("A <= B:", A <= B) // Equals - true

	// Logical operators - Boolean Operators
	AA := true
	BB := false

	//Logical operator AND
	fmt.Println("AA && BB:", AA && BB) // Equals false

	//Logical operator OR
	fmt.Println("AA || BB:", AA || BB) // Equals true
	
	//Logical operator NOT
	fmt.Println("!AA:", !AA) // Equals false
	fmt.Println("!BB:", !BB) // Equals true
	// ********************************

	// Assignment Operators
	AAA := 10
	BBB := 20

	AAA += BBB                    // A = A + B
	fmt.Println("AAA += BBB:", AAA) // Equals 30

	A -= BBB                    // A = A - B
	fmt.Println("AAA -= BBB:", AAA) // Equals 10

	// ************************************
}
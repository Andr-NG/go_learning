// Pointers store the memory address of a variable

// Declaring pointers 
// var int_pointer *int
// var float_pointer *float32

// Basic operations:
// Storting the address (&) and accessing the value via that address (dereferencing/*)


package main
import "fmt"

func main() {
	// Daclaring a variable
	var char int = 20

	// Declaring a pointer
	var charPointer *int
	charPointer = &char
	fmt.Printf("Memory address of var is %x\n", &char)
	fmt.Printf("Memory address stored in charPointer is %x\n", charPointer)
	fmt.Printf("Value of char of charPointer is %d\n", *charPointer)

	// Nil pointers
	var ptr *int
	fmt.Printf("Value of nil pointer ptr is: %x\n", ptr)

}
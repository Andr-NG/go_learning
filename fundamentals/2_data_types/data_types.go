package main

import "fmt"

func main() {


	// 1.Integers
	// signed integers store both positive and negative values
	var signed int = -500

	// unsigned integeres store only positive values
	var unsigned uint = 500

	fmt.Printf("Signed integer is %v\n", signed)
	fmt.Printf("Unsigned integer is %v\n", unsigned)

	// 2.Float
	// float32 is used for less precise calculations
	var FLOAT32 float32 = 10.6

	// float64 is used for less precise calculations
	var FLOAT64 float64 = 10.6131

	fmt.Println("FLOAT32: ", FLOAT32)
	fmt.Println("FLOAT64: ", FLOAT64)

	// 2.Boolean
	var isTrue bool = true
	var isFalse bool = true

	fmt.Println("IsTrue: ", isTrue)
	fmt.Println("IsFalse: ", isFalse)
	
	// 3.String
	var name string = "I am Golang"
	fmt.Printf("name is of type: %T\n", name)

	// 4.Complex

	// var CN1 complex128 = complex(5, 10)
	// var CN2 complex64 = complex(2, 7)

}

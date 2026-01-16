// SLICES

package main

import "fmt"



func main() {

	// Initialising a slice with 0 x5
	numbers := make([]int, 5)
	fmt.Println(numbers)

	// Subslicing
	customSlice := []int{1,2,3,4,5,6,7,8,9,10}
	subSlice := customSlice[1:4] // [2 3 4]
	subSlice2 := customSlice[3:] // [3 4 5 6 7 8 9 10]  

	// Appending
	subSlice = append(subSlice, 4, 3, 2)

	fmt.Println("subSlice2 is", subSlice2)
	fmt.Println("Appended slice is", subSlice)


	// Nil slices
	var nilSlice []int
	fmt.Println("nilSlice is", nilSlice)

	if nilSlice == nil {
		fmt.Println("nilSlice is indeed nil!")
	}
	nilSlice = append(nilSlice, 10, 10)

	if nilSlice == nil {
		fmt.Println("nilSlice is indeed nil!")
	} else {
		fmt.Println("Not nil anymore")
	}

}
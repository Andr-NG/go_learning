package main

import "fmt"

func main() {

	// result := ReverseSlice([]int{1, 2, 3, 4, 5, 6})
	// result := DeleteElement([]string{"a", "b", "c", "d"}, 2)
	// result := AppendSlice([]int{1, 2, 4, 5}, 3, 2)
	// result := FilterEvenNumbers([]int{1, 2, 3, 4, 5, 6})
	// result := DeduplicateInt([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	// result := ChunkSlice([]int{1,2,3,4,5,6,7}, 3)
	// result := RotateSlice([]int{1, 2, 3, 4, 5}, 2)
    // result := RemoveEmptyElems([]string{"go", "", "is", "", "fun"})
    // result := FlattenSlice([][]int{{1,2}, {3}, {4,5,6}})
    result := CountWord("go go rust python go rust")
	fmt.Println(result)

	// original := []int{1, 2, 3, 4, 5, 6}
	// originalPointer := &original

	// // modifying this copy will not change the original slice
	// anotherCopy := make([]int, len(original))
	// copy(anotherCopy, original)

	// // copySlice shares the same underlying array as parrent
	// copySlice := original[:6]
	// copyPointer := &copySlice

	// // Modiyfing the copy also changes the original slice
	// copySlice[4] = 7 // [1 2 3 4 7 6]
	// copySlice[5] = 8 // [1 2 3 4 7 8]
	// copySlice[1] = 8 // [1 8 3 4 7 8]
	// fmt.Printf("Original header pointer: %p\n", originalPointer)
	// fmt.Printf("Copy header pointer: %p\n", copyPointer)
	// fmt.Printf("Original data pointer: %p\n", &original[0])
	// fmt.Printf("Copy data pointer: %p\n", &copySlice[0])
	// fmt.Println("Modifying the copy slice...")
	// fmt.Println("Original slice:", original) // [1 8 3 4 7 8]
	// fmt.Println("Copy slice:", copySlice) // [1 8 3 4 7 8]
	// fmt.Println("Another copy:", anotherCopy)

}

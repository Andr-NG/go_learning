package main

// import "fmt"

func FlattenSlice(lst [][]int) []int {

    resultSlice := make([]int, 0)

	for _, val := range lst {
        resultSlice = append(resultSlice, val...)
	}
    return resultSlice
}
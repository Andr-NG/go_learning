package main

func DeduplicateInt(lst []int) []int {

	tempMap := make(map[int]int)
	resultSlice := make([]int, 0, 10)

	for _, val := range lst {
		tempMap[val]++
        if tempMap[val] == 1 {
			resultSlice = append(resultSlice, val)
		}
	}
	return resultSlice
}

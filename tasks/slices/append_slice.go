package main

import "slices"

func AppendSliceOldWay(lst []int, value, index int) []int {
	
    // 1. Add a dummy element to expand the capacity/length [1, 2, 4, 5, 0]
	lst = append(lst, 0)

    // 2. Shift elements to the right: [1, 2, 4, 4, 5]
    copy(lst[index+1:], lst[index:])

    // 3. Insert the value: [1, 2, 3, 4, 5]
    lst[index] = value    
    return lst
}


func AppendSlice(lst []int, value, index int) []int {
	
    lst = slices.Insert(lst, index, value)
    return lst
}
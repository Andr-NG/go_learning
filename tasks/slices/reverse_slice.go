// two-pointer approach, index swapping


package main

func ReverseSlice(lst []int) []int {

	lastIndex := len(lst) - 1

	i, j := 0, lastIndex
	for i < j {
		lst[i], lst[j] = lst[j], lst[i]
		i++
		j--
	}
    return lst
}

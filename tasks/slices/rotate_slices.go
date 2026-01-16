//

package main

import (
	"slices"
)

func RotateSlice(lst []int, rotate int) []int {
	tempSlice := lst[len(lst)-rotate:]

	lst = slices.Concat(tempSlice, lst[:len(lst)-rotate])
	return lst

}
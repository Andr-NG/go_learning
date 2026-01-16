// no methods to remove by index
// removing an element occurs with slicing the original slice



package main

func DeleteElement(lst []string, ind int) []string {
    lst = append(lst[:ind], lst[ind+1:]...)
    return lst
}
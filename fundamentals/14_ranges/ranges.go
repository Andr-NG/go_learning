// Range

package main

import "fmt"

func main() {
	numbers := []int{1,2,3,4,5,6,7,8,9,10}

	for ind, val := range numbers {
		fmt.Printf("Index is %d. Item is %d\n", ind, val)
	}
	var a [10]int
	fmt.Println(a)
}
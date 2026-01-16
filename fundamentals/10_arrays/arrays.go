package main

import "fmt"

func main() {
	var list [10]int

	for i := 0; i < 10; i++ {
		list[i] = i + 10
	}

	fmt.Println(list)

	for i := range list {
		fmt.Printf("Element [%d] = %d\n", i, list[i])
	}


}

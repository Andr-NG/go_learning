package main

import "fmt"

func main() {

	nums := []int{2, 3, 4, 5, 6, 7, 8, 121}
	k := 4

	windowSum := 0
	for i := range k {
		windowSum += nums[i]
	}

	maxSum := windowSum

	for right := k; right < len(nums); right++ {
		windowSum += nums[right]
		windowSum -= nums[right-k]

		if windowSum > maxSum {
			maxSum = windowSum
		}
	}

	fmt.Println(maxSum)
}

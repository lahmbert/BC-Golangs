package main

import (
	"fmt"
)

func main() {
	// nums := []int{1, 2, 3, 4, 4, 4, 3, 3, 2, 4}
	nums := []int{1, 1, 1, 2, 2, 2, 3, 3, 3}
	result := countFrequentElement(nums)

	for key, value := range result {
		fmt.Printf("%d=%d, ", key, value)
	}
	fmt.Println()
}

func countFrequentElement(nums []int) map[int]int {
	frequency := make(map[int]int)

	for _, num := range nums {
		frequency[num]++
	}

	return frequency
}

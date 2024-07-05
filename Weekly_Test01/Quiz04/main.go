package main

import (
	"fmt"
)

func main() {
	number := []int{1, 2, 2, 3, 3, 4, 4, 4}
	// words2 := "lalamama"
	result := removeDuplicate(number)
	fmt.Printf("%d\n", result)
}

func removeDuplicate(nums []int) []int {
	var result []int
	for index, value := range nums {
		same := false
		for i := 0; i < len(result); i++ {
			if value == result[i] {
				same = true
				break
			}
		}
		if !same {
			result = append(result, nums[index])
		}
	}

	return result
}

package main

import "fmt"

func main() {
	// nums := []int{1, 2, 3, 4}
	// nums := []int{1, 4, 8, 9}
	nums := []int{9, 9, 9, 9}
	result := plusOne(nums)
	fmt.Println(result)
}

func plusOne(nums []int) []int {
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < 9 {
			nums[i]++
			return nums
		} else {
			nums[i] = 0
		}
	}
	return append([]int{1}, nums...)
}

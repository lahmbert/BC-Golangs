package main

import (
	"fmt"
)

func main() {
	nums1 := []int{1, 2, 3}
	nums2 := []int{4, 5, 6}
	// nums1 := []int{9, 2, 7}
	// nums2 := []int{1, 3, 5}

	result := addDigits(nums1, nums2)

	fmt.Println(result)
}

func addDigits(nums1 []int, nums2 []int) []int {
	resultLen := len(nums2) + 1
	result := make([]int, resultLen)

	i, j := len(nums1)-1, len(nums2)-1
	k := resultLen - 1

	for i >= 0 || j >= 0 {
		sum := result[k]

		if i >= 0 {
			sum += nums1[i]
		}
		if j >= 0 {
			sum += nums2[j]
		}

		result[k] = sum % 10
		if sum >= 10 {
			result[k-1] += sum / 10
		}

		i--
		j--
		k--
	}

	if result[0] == 0 {
		result = result[1:]
	}

	return result
}

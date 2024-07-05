package main

import (
	"fmt"
	"sort"
)

func main() {
	// nums1 := []int{1, 2, 4, 7, 8}
	// nums2 := []int{2, 3, 7, 9}
	// nums1 := []int{7, 7, 3, 2, 9}
	// nums2 := []int{1, 2, 7, 4, 7, 8}
	nums1 := []int{2, 4, 6, 8}
	nums2 := []int{1, 3, 5, 7, 9}
	nums1 = removeDuplicate(nums1)
	nums2 = removeDuplicate(nums2)
	hasil := sameElement(nums1, nums2)
	fmt.Println(hasil)
}

func removeDuplicate(nums []int) []int {
	sort.Ints(nums)
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

func sameElement(nums1 []int, nums2 []int) []int {
	hasil := []int{}
	for _, v := range nums1 {
		target := false
		for _, value := range nums2 {
			if value == v {
				target = true
				break
			}
		}
		if target {
			hasil = append(hasil, v)
		}
	}
	return hasil
}

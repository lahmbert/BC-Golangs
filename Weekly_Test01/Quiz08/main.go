package main

import (
	"fmt"
)

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

func main() {
	nums1 := []int{1, 2, 3, 4, 5}
	nums2 := []int{3, 4, 5, 6, 7}
	result := sameElement(nums1, nums2)
	fmt.Println(result) // Output: [1 2]
}

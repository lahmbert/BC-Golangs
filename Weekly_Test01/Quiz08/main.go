package main

import (
	"fmt"
	"sort"
)

func main() {
	// nums1 := []int{1, 2, 3, 4, 5}
	// nums2 := []int{2, 4, 6, 7}
	nums1 := []int{2, 4, 6}
	nums2 := []int{1, 3, 5, 7}
	result1 := sliceOperation(nums1, nums2, 1)
	result2 := sliceOperation(nums1, nums2, 2)
	result3 := sliceOperation(nums1, nums2, 3)
	result4 := sliceOperation(nums1, nums2, 4)
	fmt.Println(result1)
	fmt.Println(result2)
	fmt.Println(result3)
	fmt.Println(result4)
}

func sliceOperation(nums1 []int, nums2 []int, typeOperation int) []int {
	var result []int
	switch typeOperation {
	case 1:
		hasil := []int{}
		for _, v := range nums1 {
			target := false
			for _, value := range nums2 {
				if value == v {
					target = true
					break
				}
			}
			if !target {
				hasil = append(hasil, v)
			}
			result = hasil
		}
	case 2:
		hasil := []int{}
		for _, v := range nums2 {
			target := false
			for _, value := range nums1 {
				if value == v {
					target = true
					break
				}
			}
			if !target {
				hasil = append(hasil, v)
			}
			result = hasil
		}
	case 3:
		var unionAB []int
		unionAB = append(unionAB, nums1...)
		unionAB = append(unionAB, nums2...)
		resultNotSort := append(result, unionAB...)
		sort.Ints(resultNotSort)
		result = removeDuplicate(resultNotSort)
	case 4:
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
			result = hasil
		}
	}
	return result

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

package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 2, 3, 10, 4, 6, 5, 8, 9, 7}
	sort.Ints(nums)
	result := oddBeforeEven(nums)
	fmt.Println(result)
}

func oddBeforeEven(nums []int) []int {
	even := []int{}
	odd := []int{}

	for _, v := range nums {
		if v%2 != 0 {
			even = append(even, v)
		} else {
			odd = append(odd, v)
		}
	}

	return append(even, odd...)
}

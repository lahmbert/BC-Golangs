package main

import (
	"fmt"
)

func main() {
	fruits1 := []string{"Mangga", "Apel", "Melon", "Pisang", "Sirsak", "Tomat", "Nanas", "Nangka", "Timun", "Mangga"}
	fruits2 := []string{"Bayam", "Wortel", "Kangkung", "Mangga", "Tomat", "Kembang_Kol", "Nangka", "Timun"}

	result1 := sliceFruits(fruits1, fruits2, 1)
	result2 := sliceFruits(fruits1, fruits2, 2)

	fmt.Println("OperationType = 1 (Same):", result1)
	fmt.Println("OperationType = 2 (Difference):", result2)
}

func sliceFruits(fruits1 []string, fruits2 []string, operationType int) []string {
	var result []string
	mapFruits := make(map[string]bool)

	for _, fruit := range fruits1 {
		mapFruits[fruit] = true
	}

	switch operationType {
	case 1:
		for _, fruit := range fruits2 {
			if mapFruits[fruit] {
				result = append(result, fruit)
			}
		}
	case 2:
		var result []string
		for _, v := range fruits1 {
			target1 := false
			for _, value := range fruits2 {
				if value == v {
					target1 = true
					break
				}
			}
			if !target1 {
				result = append(result, v)
			}
		}
		for _, v := range fruits2 {
			target1 := false
			for _, value := range fruits1 {
				if value == v {
					target1 = true
					break
				}
			}
			if !target1 {
				result = append(result, v)
			}
		}
		return result

	}

	return result
}

package main

import "fmt"

func main() {
	words := "bananas"
	words2 := "lalamama"
	result := removeDuplicateLetter(words)
	fmt.Printf("%s\n", result)
	result2 := removeDuplicateLetter(words2)
	fmt.Printf("%s\n", result2)
}

func removeDuplicateLetter(words string) string {
	var result string
	var sameStr string
	for _, value := range words {
		same := false
		for i := 0; i < len(sameStr); i++ {
			if rune(sameStr[i]) == value {
				same = true
				break
			}
		}
		if !same {
			sameStr += string(value)
			result += string(value)
		}
	}

	return result
}

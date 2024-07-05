package main

import (
	"fmt"
	"strings"
)

func main() {
	// word1 := "Otto"
	// word2 := "Toto"
	word1 := "Ayam"
	word2 := "Maya"
	// word1 := "Kiamat"
	// word2 := "Tamat"
	// word1 := "Kiamat"
	// word2 := "Tiamat"
	result := isAnagram(word1, word2)
	// result := sortString(strings.ToLower(word1)) == sortString(strings.ToLower(word2))
	fmt.Println(result)
	fmt.Println(sortString(word1))
	fmt.Println(sortString(word2))
}

func sortString(word string) string {
	runeWords := []rune(word)
	n := len(runeWords)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if runeWords[j] > runeWords[j+1] {
				runeWords[j], runeWords[j+1] = runeWords[j+1], runeWords[j]
			}
		}
	}
	return string(runeWords)
}

func isAnagram(word1 string, word2 string) bool {
	word1 = strings.ToLower(word1)
	word2 = strings.ToLower(word2)
	sortedWord1 := sortString(word1)
	sortedWord2 := sortString(word2)

	if len(word1) != len(word2) {
		return false
	}

	return sortedWord1 == sortedWord2
}

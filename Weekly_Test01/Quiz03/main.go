package main

import (
	"fmt"
	"strings"
)

func main() {
	word := []string{"this", "is", "a", "kanoha"}
	except := []string{"is", "a"}
	result := capitalize(word, except)
	fmt.Println(result)
}

func capitalize(words []string, except []string) []string {
	for i, word := range words {
		target := false
		for _, ex := range except {
			if word == ex {
				target = true
				break
			}
		}
		if !target {
			words[i] = strings.Title(word)
		}
	}
	return words
}

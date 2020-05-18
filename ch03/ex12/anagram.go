package main

import (
	"fmt"
	"strings"
)

func isAnagram(a, b string) bool {
	normalA := strings.ToLower(strings.ReplaceAll(a, " ", ""))
	normalB := strings.ToLower(strings.ReplaceAll(b, " ", ""))

	if len(normalA) != len(normalB) {
		return false
	}

	for i := 0; i < len(normalA); i++ {
		if char := normalA[i:i+1]; strings.Count(normalA, char) != strings.Count(normalB, char) {
			return false
		}
	}
	return true
}

func main() {
	sample := []string{
		"pat", "tap",
		"pat", "top",
		"pat", "tape",
		"roast beef", "eat for BSE",
		"New York Times", "monkeys write",
	}

	for i := 0; i < len(sample); i += 2 {
		fmt.Printf("%q, %q -> %v\n", sample[i], sample[i+1], isAnagram(sample[i], sample[i+1]))
	}
}

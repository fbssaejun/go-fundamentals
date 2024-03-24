package main

import (
	"fmt"
	"strings"
)

func countUniqueWords(messages []string) int {
	uniqueWords := make(map[string]bool)
	for _, message := range messages {
		words := strings.Fields(message)
		for _, word := range words {
			uniqueWords[word] = true
		}
	}
	return len(uniqueWords)
}

func main() {
	messages := []string{
		"Hello, world!",
		"Hello, earth!",
		"Bye, world!",
	}
	fmt.Println(countUniqueWords(messages)) // 4
}

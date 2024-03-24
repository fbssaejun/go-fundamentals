package main

import (
	"fmt"
	"strings"
)

// 1st solution
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

// 2nd solution
func countUniqueWordsSecond(messages []string) int {
	uniqueWords := make(map[string]int)
	for _, message := range messages {
		words := strings.Fields(message)
		for _, word := range words {
			uniqueWords[word]++
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
	fmt.Println(countUniqueWords(messages))       // 4
	fmt.Println(countUniqueWordsSecond(messages)) // 4
}

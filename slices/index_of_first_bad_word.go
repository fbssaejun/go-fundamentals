package main

import "fmt"

func indexOfFirstBadWord(msg []string, badWords []string) int {
	for i, word := range msg {
		for _, badWord := range badWords {
			if word == badWord {
				return i
			}
		}
	}
	return -1
}

func main() {
	msg := []string{"Hello", "world", "this", "is", "a", "frick"}
	badWords := []string{"frack", "frick"}
	fmt.Println(indexOfFirstBadWord(msg, badWords))
}

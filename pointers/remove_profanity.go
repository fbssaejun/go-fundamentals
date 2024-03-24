package main

import (
	"errors"
	"strings"
)

func removeProfanity(message *string) error {
	if message == nil {
		return errors.New("invalid input")
	}
	messageVal := *message
	messageVal = strings.ReplaceAll(messageVal, "fubb", "****")
	messageVal = strings.ReplaceAll(messageVal, "shiz", "****")
	messageVal = strings.ReplaceAll(messageVal, "witch", "*****")

	*message = messageVal
	return nil
}

func main() {
	message := "What the fubb! I can't believe you've done this you little witch! shiz!"
	removeProfanity(&message)
	println(message) // What the ****! I can't believe you've done this.
}

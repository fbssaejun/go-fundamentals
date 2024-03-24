package main

import "fmt"

func getFormattedMessages(messages []string, formatter func(string) string) []string {
	formattedMessages := []string{}
	for _, message := range messages {
		formattedMessages = append(formattedMessages, formatter(message))
	}
	return formattedMessages
}

func main() {
	messages := []string{
		"Hello, world!",
		"Hello, earth!",
		"Bye, world!",
	}
	formatter := func(message string) string {
		return fmt.Sprintf("The message is: %s", message)
	}
	fmt.Println(getFormattedMessages(messages, formatter))
	// [The message is: Hello, world! The message is: Hello, earth! The message is: Bye, world!]
}

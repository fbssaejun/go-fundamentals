package main

import (
	"fmt"
	"strings"
)

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, addTag func(msg sms) []string) []sms {
	for i, message := range messages {
		messages[i].tags = addTag(message)
	}

	return messages
}

func tagger(msg sms) []string {
	tags := []string{}
	content := strings.ToLower(msg.content)

	if strings.Contains(content, "urgent") {
		tags = append(tags, "Urgent")
	}
	if strings.Contains(content, "sale") {
		tags = append(tags, "Promo")
	}

	return tags
}

func main() {
	messages := []sms{
		{"1", "Hello, world! Urgent", []string{}},
		{"2", "Hello, earth!", []string{}},
		{"3", "Bye, world! Sale", []string{}},
	}

	// Adds tags to messages
	tagMessages(messages, tagger)

	fmt.Println(messages)
}

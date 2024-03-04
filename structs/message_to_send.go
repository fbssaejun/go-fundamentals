package main

import "fmt"

type messageToSend struct {
	phoneNumber int
	message     string
	recipient   user
	sender      user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	if mToSend.sender.name == "" {
		return false
	}
	if mToSend.recipient.name == "" {
		return false
	}
	if mToSend.sender.number == 0 {
		return false
	}
	if mToSend.recipient.number == 0 {
		return false
	}

	return true
}

func main() {
	validMessegeStruct := messageToSend{
		phoneNumber: 1234567890,
		message:     "Hello, World!",
		recipient:   user{"John Doe", 1234567890},
		sender:      user{"Jane Doe", 1110001111},
	}

	invalidMessegeStruct := messageToSend{
		phoneNumber: 1234567890,
		message:     "Hello, World!",
		recipient:   user{"", 1234567890},
		sender:      user{"Jane Doe", 0},
	}

	fmt.Println(validMessegeStruct.phoneNumber, validMessegeStruct.message)
	fmt.Println(canSendMessage(validMessegeStruct))
	fmt.Println(canSendMessage(invalidMessegeStruct))
}

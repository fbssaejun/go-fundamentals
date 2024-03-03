package main

import "fmt"

type messageToSend struct {
	phoneNumber int
	message     string
}

func main() {
	messegeStruct := messageToSend{
		phoneNumber: 1234567890,
		message:     "Hello, World!",
	}

	fmt.Println(messegeStruct.phoneNumber, messegeStruct.message)
}

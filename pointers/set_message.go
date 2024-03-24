package main

import "fmt"

func (e *email) setMessage(newMessage string) {
	e.message = newMessage
}

type email struct {
	message     string
	fromAddress string
	toAddress   string
}

func main() {
	email := email{
		message:     "Hello, world!",
		fromAddress: "Toronto, Canada",
		toAddress:   "New York, USA",
	}

	email.setMessage("Hello, earth!")
	fmt.Println(email.message)
}

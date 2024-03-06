package main

import (
	"fmt"
)

func (e email) cost() int {
	if e.isSubscribed {
		return len(e.body) * 2
	}
	return len(e.body) * 5
}

func (e email) format() string {
	var subscribedMessage string
	if e.isSubscribed {
		subscribedMessage = "Subscribed"
	} else {
		subscribedMessage = "Not Subscribed"
	}
	return fmt.Sprintf("'%s' | %s", e.body, subscribedMessage)
}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

// Implements both the expense and formatter interfaces
type email struct {
	isSubscribed bool
	body         string
}

func printInfo(e expense, f formatter) string {
	return fmt.Sprintf("Email Cost: $%v | Body: %s", e.cost(), f.format())
}

func main() {
	email1 := email{
		isSubscribed: true,
		body:         "We've been reaching out to you about your car's extended warranty",
	}
	email2 := email{
		isSubscribed: false,
		body:         "We've been reaching out to you about your car's extended warranty",
	}

	fmt.Println(printInfo(email1, email1))
	fmt.Println(printInfo(email2, email2))
}

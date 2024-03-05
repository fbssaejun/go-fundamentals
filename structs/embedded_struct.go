package main

import "fmt"

type sender struct {
	rateLimit int
	user
}

type user struct {
	name   string
	number int
}

func main() {
	sender := sender{
		rateLimit: 10,
		user: user{
			name:   "John",
			number: 123,
		},
	}

	fmt.Println(sender.name + " has a rate limit of " + fmt.Sprint(sender.rateLimit) + " messages per second.")
}

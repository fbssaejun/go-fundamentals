package main

import "fmt"

func getExpenseReport(e expense) (string, float64) {
	switch v := e.(type) {
	case email:
		return v.toAddress, v.cost()
	case sms:
		return v.toPhoneNumber, v.cost()
	default:
		return "Invalid", e.cost()
	}
}

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}

func main() {
	email1 := email{
		isSubscribed: true,
		body:         "We've been reaching out to you about your car's extended warranty",
		toAddress:    "Toronto, ON",
	}
	sms1 := sms{
		isSubscribed:  false,
		body:          "We've been reaching out to you about your car's extended warranty",
		toPhoneNumber: "555-555-5555",
	}

	invalid1 := invalid{}

	fmt.Println(getExpenseReport(email1))
	fmt.Println(getExpenseReport(sms1))
	fmt.Println(getExpenseReport(invalid1))
}

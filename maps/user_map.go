package main

import (
	"errors"
	"fmt"
)

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid sizes.")
	}

	userMap := make(map[string]user)
	for i, name := range names {
		userMap[name] = user{
			name:        name,
			phoneNumber: phoneNumbers[i],
		}
	}

	return userMap, nil
}

type user struct {
	name        string
	phoneNumber int
}

func main() {
	names := []string{"John", "Jane", "Doe"}
	phoneNumbers := []int{123456, 7891011, 12131415}

	userMap, err := getUserMap(names, phoneNumbers)
	if err != nil {
		fmt.Println(err)
		return
	}

	for name, user := range userMap {
		fmt.Printf("Name: %s, Phone Number: %d\n", name, user.phoneNumber)
	}
}

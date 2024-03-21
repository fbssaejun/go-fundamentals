package main

import (
	"errors"
	"fmt"
)

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	existingUser, ok := users[name]
	if !ok {
		return false, errors.New("not found")
	}
	if existingUser.scheduledForDeletion {
		delete(users, name)
		return true, nil
	}
	return false, nil
}

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}

func main() {
	users := map[string]user{
		"John": user{name: "John", number: 123456, scheduledForDeletion: false},
		"Jane": user{name: "Jane", number: 7891011, scheduledForDeletion: true},
		"Doe":  user{name: "Doe", number: 12131415, scheduledForDeletion: false},
	}

	deleted, err := deleteIfNecessary(users, "Jane")
	if err != nil {
		fmt.Println(err)
		return
	}
	if deleted {
		fmt.Println("User deleted.")
	} else {
		fmt.Println("User not deleted.")
	}
	fmt.Println(users)
}

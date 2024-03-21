package main

import "fmt"

func getCounts(messagedUsers []string, validUsers map[string]int) map[string]int {
	for _, messagedUser := range messagedUsers {
		if _, ok := validUsers[messagedUser]; ok {
			validUsers[messagedUser]++
		}
	}
	return validUsers
}

func main() {
	validUsers := map[string]int{
		"John": 0,
		"Jane": 0,
		"Doe":  0,
	}
	messagedUsers := []string{"John", "Jane", "Jane"}

	fmt.Println(getCounts(messagedUsers, validUsers))
}

package main

import (
	"fmt"
	"slices"
)

/*
The function should return a slice of strings containing the users suggested friends. A suggested friend is someone who is not a direct friend of the user but is a direct friend of one or more of the user's direct friends. Each suggested friend should appear only once in the slice, even if they are found through multiple direct friends.

suggestedFriends := findMutualFriends("Alice", friendships)
// suggestedFriends = [David, Eve]
*/

func findSuggestedFriends(username string, friendships map[string][]string) []string {
	suggestedFriends := []string{}
	directFriends := friendships[username]

	for _, directFriend := range directFriends {
		mutualFriends := friendships[directFriend]
		for _, name := range mutualFriends {
			if name == username {
				continue
			}
			if slices.Contains(directFriends, name) {
				continue
			}
			if slices.Contains(suggestedFriends, name) {
				continue
			}
			suggestedFriends = append(suggestedFriends, name)
		}
	}

	return suggestedFriends
}

func main() {
	friendships := map[string][]string{
		"alice":   []string{"bob", "charlie"},
		"bob":     []string{"alice", "charlie", "christine"},
		"charlie": []string{"alice", "bob", "david"},
		"david":   []string{"alice", "charlie"},
	}

	fmt.Println(findSuggestedFriends("alice", friendships))   // [charlie david]
	fmt.Println(findSuggestedFriends("bob", friendships))     // [alice charlie]
	fmt.Println(findSuggestedFriends("charlie", friendships)) // [bob david]
	fmt.Println(findSuggestedFriends("david", friendships))   // [alice charlie]
}

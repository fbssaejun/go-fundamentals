package main

import "fmt"

const (
	logDeleted  = "user deleted"
	logNotFound = "user not found"
	logAdmin    = "admin deleted"
)

type user struct {
	name   string
	number int
	admin  bool
}

func logAndDelete(users map[string]user, name string) (log string) {
	defer delete(users, name)
	user, ok := users[name]
	if !ok {
		return logNotFound
	}
	if user.admin {
		return logAdmin
	}
	return logDeleted
}

func main() {
	users := map[string]user{
		"alice": user{"Alice", 123, false},
		"bob":   user{"Bob", 456, false},
		"admin": user{"Admin", 789, true},
	}

	fmt.Println(logAndDelete(users, "alice")) // user deleted
	fmt.Println(logAndDelete(users, "admin")) // admin deleted
	fmt.Println("final users:", users)        // map[admin:{Admin 789 true} bob:{Bob 456 false}]
}

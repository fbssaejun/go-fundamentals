package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

func (info authenticationInfo) getBasicAuth() string {
	return fmt.Sprintf(
		"Authorization: Basic %s:%s",
		info.username,
		info.password,
	)
}

func main() {
	authenticationInfo := authenticationInfo{
		username: "admin",
		password: "password",
	}

	fmt.Println(authenticationInfo.getBasicAuth())
}

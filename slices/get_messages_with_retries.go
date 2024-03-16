package main

import "fmt"

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int, error) {
	return [3]string{primary, secondary, tertiary}, [3]int{len(primary), len(secondary), len(tertiary)}, nil
}

func printMessages() {
	message, cost, err := getMessageWithRetries("primary", "secondary", "tertiary")
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < len(message); i++ {
		fmt.Println(fmt.Sprintf("%s costs %d", message[i], cost[i]))
	}
}

func main() {
	printMessages()
}

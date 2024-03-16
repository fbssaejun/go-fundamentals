package main

import (
	"errors"
	"fmt"
)

const (
	planFree = "free"
	planPro  = "pro"
)

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

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	if plan == planPro {
		fmt.Println(messages[:])
		return messages[:], nil
	}
	if plan == planFree {
		fmt.Println(messages[0:2])
		return messages[0:2], nil
	}
	fmt.Println(errors.New("unsupported plan"))
	return nil, errors.New("unsupported plan")
}

func main() {
	messages := [3]string{"first message", "second message", "third message"}
	getMessageWithRetriesForPlan(planPro, messages)
	printMessages()
}

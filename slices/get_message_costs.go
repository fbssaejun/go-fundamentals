package main

import "fmt"

func getMessageCosts(messages []string) []float64 {
	costs := make([]float64, len(messages))
	for i := 0; i < len(messages); i++ {
		message := messages[i]
		cost := float64(len(message)) * 0.01

		costs[i] = cost
	}
	return costs
}

func printCosts(costs []float64) {
	for i := 0; i < len(costs); i++ {
		fmt.Println(costs[i])
	}
}

func main() {
	messages := []string{
		"first message",
		"second message",
		"third message",
		"this is the longes and most expensive message",
	}
	costs := getMessageCosts(messages)

	printCosts(costs)
}

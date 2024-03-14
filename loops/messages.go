package main

import (
	"fmt"
)

func bulkSend(numMessages int) float64 {
	totalCost := 0.0
	for i := 0; i < numMessages; i++ {
		totalCost += 1.0 + (0.01 * float64(i))
	}

	return totalCost
}

func maxMessages(thresh float64) int {
	totalCost := 0.0
	for i := 0; ; i++ {
		totalCost += 1 + (0.01 * float64(i))
		if totalCost > thresh {
			return i
		}
	}
}

func main() {
	numMessages := 10
	fmt.Printf("Cost to send %d messages: $%.2f\n", numMessages, bulkSend(numMessages))

	threshold := 20.0
	fmt.Printf("Max messages for $%.2f: %d\n", threshold, maxMessages(threshold))
}

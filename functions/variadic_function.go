package main

import "fmt"

// ...int means it takes in a variable number of integers
func sum(nums ...int) int {
	total := 0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	return total
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1, 2))
	fmt.Println(sum(1))
	fmt.Println(sum())
}

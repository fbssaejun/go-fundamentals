package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	add := adder()
	// it will add 1 to the "sum" variable defined in the adder() function
	fmt.Println(add(1)) // 1
	fmt.Println(add(2)) // 3
	fmt.Println(add(3)) // 6
}

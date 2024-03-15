package main

import (
	"fmt"
	"math"
)

func printPrimes(max int) {
	if max >= 2 {
		fmt.Println(2)
	}
	for n := 3; n <= max; n += 2 {
		if isPrime(n) {
			fmt.Println(n)
		}
	}
}

func isPrime(n int) bool {
	sqrt := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrt; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	max := 100
	printPrimes(max)
}

package main

import "fmt"

func getMonthlyPrice(tier string) string {
	amount := 0

	if tier == "basic" {
		amount = 100 * 100
	} else if tier == "premium" {
		amount = 150 * 100
	} else if tier == "enterprise" {
		amount = 500 * 100
	} else {
		amount = 0 * 100
	}

	amountStatement := fmt.Sprintf("The monthly price for the %s tier is: %d pennies", tier, amount)
	return amountStatement

}

func main() {
	test("basic")
	test("premium")
	test("enterprise")
	test("free")
}

func test(tier string) {
	fmt.Println(getMonthlyPrice(tier))
}

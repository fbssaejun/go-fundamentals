package main

import "fmt"

func getProductMessage(tier string) string {
	// Ignoring the last returned variable from getProductInfo, as it's not used
	quantityMsg, priceMsg, _ := getProductInfo(tier)
	return "You get " + quantityMsg + " for " + priceMsg + "."
}

func getProductInfo(tier string) (string, string, string) {
	if tier == "basic" {
		return "1,000 texts per month", "$30 per month", "most popular"
	} else if tier == "premium" {
		return "50,000 texts per month", "$60 per month", "best value"
	} else if tier == "enterprise" {
		return "unlimited texts per month", "$100 per month", "customizable"
	} else {
		return "10 texts per month", "$0 per month", "free of charge"
	}
}

func main() {
	fmt.Println(getProductMessage("basic"))
	fmt.Println(getProductMessage("premium"))
	fmt.Println(getProductMessage("enterprise"))
	fmt.Println(getProductMessage("free"))
}

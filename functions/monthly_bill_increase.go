package main

import "fmt"

func monthlyBillIncrease(costPerSend, numLastMonth, numThisMonth int) int {
	var lastMonthBill int
	var thisMonthBill int
	lastMonthBill = getBillForMonth(costPerSend, numLastMonth)
	thisMonthBill = getBillForMonth(costPerSend, numThisMonth)
	return thisMonthBill - lastMonthBill
}

func getBillForMonth(costPerSend, messagesSent int) int {
	return costPerSend * messagesSent
}

func main() {
	fmt.Println(monthlyBillIncrease(1, 100, 100))
	fmt.Println(monthlyBillIncrease(1, 100, 200))
	fmt.Println(monthlyBillIncrease(1, 200, 100))
	fmt.Println(monthlyBillIncrease(1, 200, 1000))
}

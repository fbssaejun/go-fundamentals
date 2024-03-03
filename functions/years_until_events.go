package main

import "fmt"

func yearsUntilEvents(age int) (
	yearsUntilAdult,
	yearsUntilDrinking,
	yearsUntilCarRental int,
) {
	yearsUntilAdult = 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}
	yearsUntilDrinking = 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}
	yearsUntilCarRental = 25 - age
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}

	// Naked return statement returns the named return variables
	// This is not recommended for readability, but is used here for brevity and to demonstrate the syntax
	return

	//Above is the same as:
	// return yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental
}

func main() {
	fmt.Println(yearsUntilEvents(15))
	fmt.Println(yearsUntilEvents(20))
}

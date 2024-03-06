package main

import (
	"fmt"
)

type divideError struct {
	dividend float64
}

func (devideError divideError) Error() string {
	return fmt.Sprintf("can not divide %.2f by zero", devideError.dividend)
}

func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, divideError{dividend: dividend}
	}
	return dividend / divisor, nil
}

func main() {
	result, err := divide(5, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
	result2, err := divide(5, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result2)
	}
}

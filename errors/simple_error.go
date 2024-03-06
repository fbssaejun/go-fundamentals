package main

import (
	"errors"
	"fmt"
)

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0.0, errors.New("no dividing by 0")
	}
	return x / y, nil
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

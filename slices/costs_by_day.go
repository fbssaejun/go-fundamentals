package main

import "fmt"

type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) []float64 {
	costsByDay := []float64{}
	for _, cost := range costs {
		for cost.day >= len(costsByDay) {
			costsByDay = append(costsByDay, 0.0)
		}
		costsByDay[cost.day] += cost.value
	}
	return costsByDay
}

func main() {
	costs := []cost{
		{0, 1.0},
		{1, 2.0},
		{1, 3.0},
		{2, 3.0},
		{3, 4.0},
		{4, 5.0},
		{6, 7.0},
		{9, 10.0},
	}
	costsByDay := getCostsByDay(costs)
	fmt.Println(costsByDay)
}

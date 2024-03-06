package main

import "fmt"

type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

func (c contractor) getName() string {
	return c.name
}

func (c contractor) getSalary() int {
	return c.hourlyPay * c.hoursPerYear
}

type fullTime struct {
	name   string
	salary int
}

func (ft fullTime) getSalary() int {
	return ft.salary
}

func (ft fullTime) getName() string {
	return ft.name
}

func printEmployeeInfo(e employee) string {
	return fmt.Sprintf("%s makes $%v per year", e.getName(), e.getSalary())
}

func main() {
	contractor := contractor{
		name:         "John Doe",
		hourlyPay:    50,
		hoursPerYear: 2080,
	}

	fullTime := fullTime{
		name:   "Jack Black",
		salary: 100000,
	}

	fmt.Println(printEmployeeInfo(contractor))
	fmt.Println(printEmployeeInfo(fullTime))
}

package main

import (
	"fmt"

	"github.com/headfirstgo/magazine"
)

type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
}

type Employee struct {
	Name   string
	Salary float64
}

func main() {
	var employee magazine.Employee
	employee.Name = "Joy Carr"
	employee.Salary = 60000
	fmt.Println(employee.Name)
	fmt.Println(employee.Salary)
}

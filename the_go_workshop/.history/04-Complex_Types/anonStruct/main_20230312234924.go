package main

import (
	"fmt"
	"github.com/headfirstgo/magazine"
)

type Subscriber struct {
	Name		string
	Rate   float64
	Acti   		bool
} 

type Employee struct {
	Name		string
	Sala   		float64
} 

func main() {
	var employee magazine.Employee
	employee.Name = "Joy Carr"
	employee.Salary = 60000
	fmt.Println(employee.Name)
	fmt.Println(employee.Salary)
}

/*
This exercise will help you internalize pass by value/reference and pointer semantics in Go

we will build the same calculator in two different ways

[warmup]

1. create a blank CalculatorA struct - CalculatorA is the name of the struct and there is nothing inside the struct

2. CalculatorA will have 4 functions add, sub, mult, and divide - all these functions will take in 2 integers and return 1 integer as output --- all these functions should be receiver functions i.e. they should look like func (c CalculatorA) myFunctionName(inputnumamOne int, inputnumamTwo string) outputnumamThree string
*/
package calcA

import "fmt"

var prt = fmt.Println

type CalculatorA struct {
}

// var CalcA CalculatorA

func (c CalculatorA) Add(num1 int, num2 int) (result int) {
	result = num1 + num2
	return result
}

func (c CalculatorA) Sub(num1 int, num2 int) (result int) {
	result = num1 - num2
	return result
}

func (c CalculatorA) Mult(num1 int, num2 int) (result int) {
	result = num1 * num2
	return result
}

func (c CalculatorA) Divide(num1 int, num2 int) (result int) {
	if num2 == 0 {
		prt("impossible to divide by zero")
		return 0
	} else {
		result = num1 / num2
		return result
	}
}

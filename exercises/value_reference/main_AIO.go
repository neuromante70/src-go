/*
This exercise will help you internalize pass by value/reference and pointer semantics in Go

we will build the same calculator in two different ways

[warmup]

1. create a blank CalculatorA struct - CalculatorA is the name of the struct and there is nothing inside the struct

2. CalculatorA will have 4 functions add, sub, mult, and divide - all these functions will take in 2 integers and return 1 integer as output --- all these functions should be receiver functions i.e. they should look like func (c CalculatorA) myFunctionName(inputnumamOne int, inputnumamTwo string) outputnumamThree string

Now we will do the same thing as above but a bit differently

3. Create a new blank CalculatorB struct
4. Calculator will also have these 4 functions add/sub/mult/divide but you have to follow the following two restrictions when implementing these functions
4a. There can be no return keyword used - and none of the functions will declare an output/return in the signature - you need to figure out a different way to send the response/answer back
4b. the two integers that come in as input - should be 0'd out - so if the input is 4, 5 - after any of the add/mult/sub/divide functions are called they should now be 0 [even on the caller side]


Put all your code in a different calculator/math package and call them from the main package main function
*/

package main

import (
	"fmt"
	//user "BASICS/exercises/structAndMap2/models"
	// calcA "BASICS/exercises/value_reference/calculatorA"
	// calcB "BASICS/exercises/value_reference/calculatorB"
	// calcA "BASICS/exercises/value_reference/calculatorA"
	// calcB "BASICS/exercises/value_reference/calculatorB"
)

var prt = fmt.Println

type CalculatorB struct {
}

// var CalcB CalculatorB

func (c CalculatorB) Add(num1 int, num2 int, result *int) {
	*result = num1 + num2
}

func (c CalculatorB) Sub(num1 int, num2 int, result *int) {
	*result = num1 - num2
}

func (c CalculatorB) Mult(num1 int, num2 int, result *int) {
	*result = (num1) * (num2)
}

func (c CalculatorB) Divide(num1 int, num2 int, result *int) {
	if num2 == 0 {
		prt("impossible to divide by zero")
		return
	}
	*result = num1 / num2

}

func main() {
	cB := CalculatorB{}
	result := 0
	cB.Add(5, 3, &result) // 8
	prt(result)
	cB.Sub(5, 9, &result) // -4
	prt(result)
	cB.Mult(22, 33, &result) // 726
	prt(result)
	cB.Divide(4, 0, &result) // impossible to divide by zero
	prt(result)
	cB.Divide(1232, 6, &result) // 205
	prt(result)
}

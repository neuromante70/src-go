/*
This exercise will help you internalize pass by value/reference and pointer semantics in Go

Now we will do the same thing as above but a bit differently

3. Create a new blank CalculatorB struct
4. Calculator will also have these 4 functions add/sub/mult/divide but you have to follow the following two restrictions when implementing these functions
4a. There can be no return keyword used - and none of the functions will declare an output/return in the signature - you need to figure out a different way to send the response/answer back
4b. the two integers that come in as input - should be 0'd out - so if the input is 4, 5 - after any of the add/mult/sub/divide functions are called they should now be 0 [even on the caller side]


Put all your code in a different calculator/math package and call them from the main package main function
*/

package calcB

import (
	"fmt"
)

type CalculatorB struct {
}

// var CalcB CalculatorB

func (c CalculatorB) Add(num1 *int, num2 *int, result *int) {
	*result = *num1 + *num2
	*num1 = 0
	*num2 = 0
}

func (c CalculatorB) Sub(num1 int, num2 int, result *int) {
	*result = num1 - num2
	*num1 = 0
	*num2 = 0
}

func (c CalculatorB) Mult(num1 int, num2 int, result *int) {
	*result = num1 * num2
	*num1 = 0
	*num2 = 0
}

func (c CalculatorB) Divide(num1 int, num2 int, result *int) {
	if num2 == 0 {
		fmt.Println("impossible to divide by zero")
		*result = 0
		*num1 = 0
		*num2 = 0
	} else {
		*result = num1 / num2
		*num1 = 0
		*num2 = 0
	}
}

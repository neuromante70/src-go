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
	calcA "BASICS/exercises/value_reference/calculatorA"	// /home/fmontecchio/go/src is the prefix omitted
	calcB "BASICS/exercises/value_reference/calculatorB"
)

var prt = fmt.Println

func main() {
	cA := calcA.CalculatorA{}
	cB := calcB.CalculatorB{}

	prt("calculation with CalculatorA")
	prt(cA.Add(5, 3))     // 8
	prt(cA.Sub(5, 9))     // -4
	prt(cA.Mult(22, 33))  // 726
	prt(cA.Divide(4, 0))  // impossible to divide by zero
	prt(cA.Divide(12, 6)) // 2

	prt("calculation with CalculatorB")
	result := 0
	num1 := 5
	num2 := 3
	cB.Add(&num1, &num2, &result) // 8
	prt("addition is", result)
	prt("Now num1 and num2 are", num1, num2)
	cB.Sub(5, 9, &result)
	prt("substracton is", result) // -4
	cB.Mult(22, 33, &result)
	prt("moltiplication is", result) // 726
	cB.Divide(4, 0, &result)
	prt("division is", result) // impossible to divide by zero
	cB.Divide(12, 6, &result)
	prt("division is", result) // 2

}

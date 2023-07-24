package calculator

import (
	"fmt"
)

var prt = fmt.Println

/*
1. Create a MathRequest struct - this will have two integers and an operation string - this is how we will model all the add/mult/sub/divide requests that come into the calculator.
So an add function call will create a math request where the two ints are the two inputs and the operation is "add"
*/

type MathRequest struct {
	Num1     int    `json:"Num1"`
	Operator string `json:"Operator"`
	Num2     int    `json:"Num2"`
}

/*
2. In the CalculatorA struct create a slice of math requests named history - this is where we will append all math requests that are created in the different functions - so this will be our list of history
*/

type Calculator struct {
	History []MathRequest
}

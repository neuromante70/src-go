/*
Json allows us to serialize our data to use somewhere else
Serialization is also how we can 'save' our work to be used later - think saving in games and then reloading it so you dont have to start from the start

In this exercise we will take CalculatorA from the last exercise with the add, mult, divide, and sub functions and add some new functionality to it

You might have noticed that Calculators also have some semblance of history of the operations performed - so that is what we will add

1. Create a MathRequest struct - this will have two integers and an operation string - this is how we will model all the add/mult/sub/divide requests that come into the Calculator. So an add function call will create a math request where the two ints are the two inputs and the operation is "add"

2. In the CalculatorA struct create a slice of math requests named history - this is where we will append all math requests that are created in the different functions - so this will be our list of history

3. create an encodeHistory function this will loop over all the mathRequests in the history slice inside the CalculatorA struct and convert all of them into json strings and return a list of strings. You can access the history inside the struct if the function is a receiver using the dot syntax

so something like this

json.Marshal is how you convert a struct into a []byte - and then you can convert []byte into string
func (c Calculator) encodeHistory() []string {
    for _, mr := range c.history {
    }
}
4. create a decodeHistory function that will take in a list of json strings as input and convert them into mathRequests and save them as the history inside the Calculator slice . At the start of this function clear out whatever was in the slice previously

json.Unmarshal is how you convert a []byte json into a struct passed in as a pointer
*/

package main

import (
	calculator "BASICS/exercises/json/calculator"
	"fmt"
)

var prt = fmt.Println

func main() {
	prt("calculation with Calculator:")
	result := calculator.Calculator..Add(5, 3)
	prt("5 + 3 =", result)      // 8
	prt("5 - 9 =", calculator.Calculator.Sub(5, 9))      // -4
	prt("22 * 33 =", calculator.Calculator.Mult(22, 33)) // 726
	prt("4 / 0 =", calculator.Calculator.Divide(4, 0))   // impossible to divide by zero
	prt("12 / 6 =", calculator.Calculator.Divide(12, 6)) // 2
	prt()
	_ = calculator.Calculator.EncodeHistory()
	calculator.Calculator.SaveToFile(calculator.Calculator.EncodeHistory(), "test.json")
	// opCalc.DecodeHistory()
	History2 := []byte(`[{"Num1":6,"Operator":"Add","Num2":8},{"Num1":2,"Operator":"Sub","Num2":7},{"Num1":8,"Operator":"Mult","Num2":8},{"Num1":33,"Operator":"Div","Num2":11}]`)
	calculator.Calculator.SaveToFile(History2, "test2.json")
	calculator.Calculator.DecodeHistory("test2.json")
	calculator.Calculator.AddToHistory("test2.json")
	calculator.Calculator.PrintHistory()
}

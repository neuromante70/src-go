/*
Json allows us to serialize our data to use somewhere else
Serialization is also how we can 'save' our work to be used later - think saving in games and then reloading it so you dont have to start from the start

In this exercise we will take CalculatorA from the last exercise with the add, mult, divide, and sub functions and add some new functionality to it

You might have noticed that Calculators also have some semblance of history of the operations performed - so that is what we will add

1. Create a mathRequest struct - this will have two integers and an operation string - this is how we will model all the add/mult/sub/divide requests that come into the Calculator. So an add function call will create a math request where the two ints are the two inputs and the operation is "add"

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
	"encoding/json"
	"fmt"
	"os"
)

var prt = fmt.Println

/*
1. Create a MathRequest struct - this will have two integers and an operation string - this is how we will model all the add/mult/sub/divide requests that come into the calculator.
So an add function call will create a math request where the two ints are the two inputs and the operation is "add"
*/

type mathRequest struct {
	num1     int    `json:"num1"`
	operator string `json:"Operator"`
	num2     int    `json:"num2"`
}

/*
2. In the CalculatorA struct create a slice of math requests named history - this is where we will append all math requests that are created in the different functions - so this will be our list of history
*/

type calculator struct {
	history []mathRequest
}

func (c calculator) add(num1 int, num2 int) int {
	op := mathRequest{num1, "add", num2}
	c.history = append(c.history, op)
	return num1 + num2
}

func (c calculator) sub(num1 int, num2 int) int {
	op := mathRequest{num1, "sub", num2}
	c.history = append(c.history, op)
	return num1 - num2
}

func (c calculator) mult(num1 int, num2 int) int {
	op := mathRequest{num1, "mult", num2}
	c.history = append(c.history, op)
	return num1 * num2
}

func (c calculator) divide(num1 int, num2 int) int {
	if num2 == 0 {
		prt("impossible to divide by zero")
		return 0
	}
	op := mathRequest{num1, "Div", num2}
	c.history = append(c.history, op)
	return num1 / num2
}

/*
3. create an encodeHistory function this will loop over all the mathRequests in the history slice inside the calculatorA struct and convert all of them into json strings and return a list of strings. You can access the history inside the struct if the function is a receiver using the dot syntax. So something like this:
func (c calculator) encodeHistory() []string {
    for _, mr := range c.history {
    }
}
*/

func check(e error) (ok bool) {
	if e != nil {
		fmt.Printf("Error: %s", e.Error())
		return false
	}
	return true
}

func (c calculator) EncodeHistory() (eh []byte) {
	eh, err := json.Marshal(c.History)
	if !check(err) {
		return nil
	}
	// prt("this is returned from json converter:", string(j))
	return eh
}

func (c calculator) SaveToFile(eh []byte, fn string) {
	f, err := os.Create(fn)
	if !check(err) {
		return
	}
	n2, err := f.Write(eh) // I should check if the file already exists but...
	if !check(err) {
		f.Close()
	}
	prt(n2, "bytes written successfully")
	err = f.Close()
	if !check(err) {
		return
	}
	f2, err := os.ReadFile(fn)
	if !check(err) {
		return
	}
	prt("this is ", fn, " file: ", string(f2))
}

/*
4. create a decodeHistory function that will take in a list of json strings as input and convert them into mathRequests and save them as the history inside the calculator slice . At the start of this function clear out whatever was in the slice previously
json.Unmarshal is how you convert a []byte json into a struct passed in as a pointer
*/

func (c calculator) DecodeHistory(fn string) (f []byte) {
	f, err := os.ReadFile(fn)
	if !check(err) {
		return
	}
	// var data []interface{}
	// err = json.Unmarshal(f, &data)
	mr := make([]mathRequest, 10)
	err = json.Unmarshal(f, &mr)
	if !check(err) {
		return
	}
	return f
}

func (c calculator) addToHistory(fn string) {
	mr := make([]mathRequest, 10)
	f, err := os.ReadFile(fn)
	if !check(err) {
		return
	}
	err = json.Unmarshal(f, &mr)
	if !check(err) {
		return
	}
	// for i := 0; i < len(mr)-1; i++ {
	// 	MyCalc.History[i].num1 = mr[i].num1
	// 	MyCalc.History[i].Operator = mr[i].Operator
	// 	MyCalc.History[i].num2 = mr[i].num2
	// }
	c.History = nil
	for _, mRequest := range mr {
		c.History = append(c.History, mathRequest(mRequest))
	}
	// for i := 0; i < len(mr); i++ {
	// 	op := mathRequest{mr[i].Num1, mr[i].Operator, mr[i].num2}
	// 	MyCalc.History = append(MyCalc.History, op)
	// }
	prt("this is the last history:", c.PrintHistory())
}

// ancillary function
func (c calculator) PrintHistory() (history []string) {
	for _, v := range c.History {
		history = append(history, fmt.Sprintf("%+v", v))
	}
	return history
}

func main() {
	prt("calculation with calculator:")
	prt("5 + 3 =", calc.add(5, 3))      // 8
	prt("5 - 9 =", calc.sub(5, 9))      // -4
	prt("22 * 33 =", calc.mult(22, 33)) // 726
	prt("4 / 0 =", calc.divide(4, 0))   // impossible to divide by zero
	prt("12 / 6 =", calc.divide(12, 6)) // 2
	prt()
	_ = calc.EncodeHistory()
	calc.SaveToFile(calc.EncodeHistory(), "test.json")
	// opCalc.DecodeHistory()
	History2 := []byte(`[{"num1":6,"Operator":"add","num2":8},{"num1":2,"Operator":"sub","num2":7},{"num1":8,"Operator":"mult","num2":8},{"num1":33,"Operator":"Div","num2":11}]`)
	calc.SaveToFile(History2, "test2.json")
	calc.DecodeHistory("test2.json")
	calc.addToHistory("test2.json")
	calc.PrintHistory()
}

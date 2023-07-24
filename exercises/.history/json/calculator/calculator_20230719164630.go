package calc

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

// var MyCalc Calculator

func (c Calculator) Add(num1 int, num2 int) int {
	op := MathRequest{num1, "Add", num2}
	c.History = append(c.History.History, op)
	return num1 + num2
}

func (c Calculator) Sub(num1 int, num2 int) int {
	op := MathRequest{num1, "Sub", num2}
	MyCalc.History = append(MyCalc.History, op)
	return num1 - num2
}

func (c Calculator) Mult(num1 int, num2 int) int {
	op := MathRequest{num1, "Mult", num2}
	MyCalc.History = append(MyCalc.History, op)
	return num1 * num2
}

func (c Calculator) Divide(num1 int, num2 int) int {
	if num2 == 0 {
		prt("impossible to divide by zero")
		return 0
	}
	op := MathRequest{num1, "Div", num2}
	MyCalc.History = append(MyCalc.History, op)
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

func (c Calculator) EncodeHistory() (eh []byte) {
	eh, err := json.Marshal(MyCalc.History)
	if !check(err) {
		return nil
	}
	// prt("this is returned from json converter:", string(j))
	return eh
}

func (c Calculator) SaveToFile(eh []byte, fn string) {
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

func (c Calculator) DecodeHistory(fn string) (f []byte) {
	f, err := os.ReadFile(fn)
	if !check(err) {
		return
	}
	// var data []interface{}
	// err = json.Unmarshal(f, &data)
	mr := make([]MathRequest, 10)
	err = json.Unmarshal(f, &mr)
	if !check(err) {
		return
	}
	return f
}

func (c Calculator) AddToHistory(fn string) {
	mr := make([]MathRequest, 10)
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
	// 	MyCalc.History[i].Num2 = mr[i].Num2
	// }
	MyCalc.History = nil
	for _, mRequest := range mr {
		MyCalc.History = append(MyCalc.History, MathRequest(mRequest))
	}
	// for i := 0; i < len(mr); i++ {
	// 	op := MathRequest{mr[i].Num1, mr[i].Operator, mr[i].Num2}
	// 	MyCalc.History = append(MyCalc.History, op)
	// }
	prt("this is the last history:", c.PrintHistory())
}

// ancillary function
func (c Calculator) PrintHistory() (history []string) {
	for _, v := range MyCalc.History {
		history = append(history, fmt.Sprintf("%+v", v))
	}
	return history
}

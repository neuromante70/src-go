package main

/* Activity 4.06: Type Checker
Create a function that accepts a value of any type. The function returns a string with the name of the type:
• For int, int32, and int64, it returns int.
• For all floats, it returns float.
• For a string, it returns string.
• For a bool, it returns bool.
• For anything else, it returns unknown.
• Loop all the data by passing each one to your function.
*/

import (
	"errors"
	"fmt"
)

func returnType(v interface{}) (string, error) {
	switch t := v.(type) {
	case string:
		return fmt.Sprintf("Your variable is a %T\n", t), nil
	case bool:
		//if t {
			return fmt.Sprintf("Your variable is %T\n", t), nil
		//}
		return fmt.Sprintf("Your variable is %T\n", t), nil
	case float32, float64:
		return fmt.Sprintf("Your variable is a %T\n", t), nil
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("Your variable is a %T\n", t), nil
	default:
		return "", errors.New("unsupported type passed")
	}
}

func main() {
	res, _ := returnType(-5)
	fmt.Println("-5 :", res)
	res, _ = returnType(5)
	fmt.Println("5 :", res)
	res, _ = returnType("yum")
	fmt.Println("yum:", res)
	res, _ = returnType(false)
	fmt.Println("bool:", res)
	res, _ = returnType(float32(3.14))
	fmt.Println("3.14 :", res)
	res, _ = returnType(struct{}{})
	fmt.Println("struct :", res)
}

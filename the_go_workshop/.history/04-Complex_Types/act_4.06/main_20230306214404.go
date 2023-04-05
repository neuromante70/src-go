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

func doubler(v interface{}) (string, error) {
	switch t := v.(type) {
	case string:
		return fmt.Sprintf("Your variable is a string: %T\n", t), nil
	case bool:
		if t {
			return fmt.Sprintf("Your variable is bool: %T\n", t), nil
		}
		return fmt.Sprintf("Your variable is bool: %T\n", t), nil
	case float32, float64:
		// if f, ok := t.(float64); ok {
			return fmt.Sprintf("Your variable is a float: %T\n", t), nil
		/ }
		// return fmt.Sprint(t.(float32) * 2), nil
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("Your variable is an int: %T\n", t), nil
		// return fmt.Sprint(t * 2), nil
	// case int8:
	// 	return fmt.Sprint(t * 2), nil
	// case int16:
	// 	return fmt.Sprint(t * 2), nil
	// case int32:
	// 	return fmt.Sprint(t * 2), nil
	// case int64:
	// 	return fmt.Sprint(t * 2), nil
	// case uint:
	// 	return fmt.Sprint(t * 2), nil
	// case uint8:
	// 	return fmt.Sprint(t * 2), nil
	// case uint16:
	// 	return fmt.Sprint(t * 2), nil
	// case uint32:
	// 	return fmt.Sprint(t * 2), nil
	// case uint64:
	// 	return fmt.Sprint(t * 2), nil
	default:
		return "", errors.New("unsupported type passed")
	}
}

func main() {
	res, _ := doubler(-5)
	fmt.Println("-5 :", res)
	res, _ = doubler(5)
	fmt.Println("5 :", res)
	res, _ = doubler("yum")
	fmt.Println("yum:", res)
	res, _ = doubler(true)
	fmt.Println("true:", res)
	res, _ = doubler(float32(3.14))
	fmt.Println("3.14 :", res)
}

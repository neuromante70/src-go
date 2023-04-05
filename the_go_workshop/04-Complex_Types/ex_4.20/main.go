package main

// Exercise 4.20: Numeric Type Conversion
import (
	"fmt"
	"math"
)

func convert() string {
	var i8 int8 = math.MaxInt8
	i := 28
	f64 := 3.14
	//  Here, we'll convert from a smaller int type into a larger int type. This is always a safe operation
	m := fmt.Sprintf("int8 = %v > in64 = %v\n", i8, int64(i8))
	// Now, we'll convert from an int that's 1 above int8's maximum size. This will cause an overflow to int8's minimum size:
	m += fmt.Sprintf("int   = %v > in8   =%v\n", i, int8(i))
	// conversion out int8 into a float64. No overflow
	m += fmt.Sprintf("int8 = %v > float32 = %v\n", i8, float64(i8))
	// conversion a float into an int. Decimal data is lost but the whole number is kept as is
	m += fmt.Sprintf("float64 = %v > int = %v\n", f64, int(f64))
	return m
}

func main() {
	fmt.Print(convert())
}

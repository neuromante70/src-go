package main

import (
	"fmt"
	"math"
)

func main() {
	// integer max
	fmt.Printf("max int64   = %+v\n", math.MaxInt64)
	fmt.Printf("max int32   = %+v\n", math.MaxInt32)
	fmt.Printf("max int16   = %+v\n", math.MaxInt16)

	// integer min
	fmt.Printf("min int64   = %+v\n", math.MinInt64)
	fmt.Printf("min int32   = %+v\n", math.MinInt32)

	fmt.Printf("max float64 = %+v\n", math.MaxFloat64)
	fmt.Printf("max float32 = %+v\n", math.MaxFloat32)

	// etc you can see more int the `math`package
}

package main

// Working with Big Numbers, package math/big
import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	intA := math.MaxInt64
	intA += 1

	bigA := big.NewInt(math.MaxInt64)
	bigA.Add(bigA, big.NewInt(1))

	fmt.Println("MaxInt64: ", math.MaxInt64)
	fmt.Println("Int     :", intA)
	fmt.Printf("Big Int :", bigA.String())

	/*
	   Output:
	   MaxInt64:  9.223.372.036.854.775.807
	   Int     : -9.223.372.036.854.775.808
	   Big Int : 9.223.372.036.854.775.808
	*/
}

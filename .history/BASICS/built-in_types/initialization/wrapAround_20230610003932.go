package main

import (
	"fmt"
	"math"
)

func main() {
	var maxUint8 uint8 = math.MaxUint8	//uint8 or byte: range 0 to 255 (MaxUint8 constant from math package)
	fmt.Println("value", maxUint8)
	//  var maxOtherUint8 uint8 = 256	//i can't do this: this is an overflow compiler error
	//  fmt.Println("value", maxOtherUint8)
	maxUint8
}

package main

import (
	"fmt"
)

func main() {
	// var a int8 = 128	// overflow error because int8 maximum value is 127
	// fmt.Println(a)

	// Wraparound
	var a int8 = 125
	var b uint8 = 253

	for i := 0; i < 5; i++ {
		a++
		b++
		fmt.Println(i, ")", "int8", a, "uint8", b)
	}
	/* Output:
	   0 ) int8 126 uint8 254
	   1 ) int8 127 uint8 255
	   Here wraparound!
	   2 ) int8 -128 uint8 0
	   3 ) int8 -127 uint8 1
	   4 ) int8 -126 uint8 2
	*/
}

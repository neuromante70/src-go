package main

import (
	"fmt"
)

/*
FOR-RANGE LOOPS
Only iterate on:
1) strings
2) arrays and slices
3) maps
4) channels
*/

func main() {
	evenVals := []int{2, 4, 6, 8, 10, 12}
	// for index, value := intervallo slice {
	for i, v := range evenVals {
		fmt.Println("index:", i, "value:", v)
	}

	// The index or value can be omitted; replaced by the underscore.

	for _, v := range evenVals {
		fmt.Println("only range values:", v)
	}

	for i := range evenVals {	// Go 
		fmt.Println("only range indexes:", i)
	}

	// The most common reason for iterating over the key is when a map is being used as a set.
	uniqueNames := map[string]bool{"Fred": true, "Raul": true, "Wilma": false}
	for k := range uniqueNames {
		fmt.Println(k)
	}

	// The for-range value is a copy
	for _, v := range evenVals {
		v *= 2
		fmt.Println(v)
	}
	fmt.Println(evenVals) // print an unmodified version of the range

	// solution: use the index to modify the elements of the slice

	for i, v := range evenVals {
		evenVals[i] = v * 2
	}
	fmt.Println(evenVals) // print an unmodified version of the range 

}
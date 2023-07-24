package main

import "fmt"

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
	for i, v := range evenVals {
		fmt.Println("index:", i, "value:", v)
	}

}
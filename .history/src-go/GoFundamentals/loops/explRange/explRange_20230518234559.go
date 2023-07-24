package main

import "fmt"

/*
FOR-RANGE LOOPS
Only on:
*/

func main() {
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		fmt.Println("index:", i, "value:", v)
	}

}

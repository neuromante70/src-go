package main

import "fmt"
/*
for statemente for loop:
1) a complete, C-style for
2) a condition
*/


func main() {
	/* Complete for Statement, C-Style
	for {[initialization :=]; [comparison = execution if true]; [increment]
		for body
	}
	*/
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}


}

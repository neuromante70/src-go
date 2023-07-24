package main

import "fmt"

/*
for statemente for loop:
1) a complete, C-style for
2) a condition-only for (while style)
3) an infite for
   continue & break
4) for-range
*/

func main() {
	/* 1) Complete for Statement, C-Style
	for {[initialization :=]; [comparison = execution if true]; [increment]
		for body
	}
	*/
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	/* 2) a condition-only for (while style)
	i := 0
	for {[comparison = execution if true]
		for body: [increment]
	}
	*/
	n := 1
	for n < 100 {
		fmt.Println(n)
		n *= 2
	}

	/* 3) an infite for
	for {
		for body
	}
	*/

	for {
		fmt.Println("Hello")
	}

	/* do-while replacement (assure code is executed almost for 1 time)
	in Java/C:
	do {
		do body
	} while ([comparison = execution if true]) // specifies how to stay in for loop

	in Go:
	for {
		for body
		if ![comparison] {	// ! to negate the condition: specifies how to exit the loop
			break			// exit the loop! :-)
		}
	}
	*/

	// using continue to make code cleaner; else in 

	/* 4) for-range

	 */
}

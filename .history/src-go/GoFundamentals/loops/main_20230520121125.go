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
	m := 1
	for m < 100 {
		fmt.Println(m)
		m *= 2
	}

	/* 3) an infite for
	for {
		for body
	}
	*/

	for {
		fmt.Println("Infinite Hello")
		break
	}

	/* do-while replacement (assure code is executed almost for 1 time)
	in Java/C:
	do {
		do body
	} while ([comparison = execution if true]) // specifies how to stay in for loop

	in Go:
	for {
			index incrementation
			for body
		if ![comparison] {	// ! to negate the condition: specifies how to exit the loop
			break			// exit the loop! :-)
		}
		
	}
	*/
	for {
		if !m
	}

	// using continue to make code cleaner; else in if statemente is not necessary
	// FizzBuzz problem: Print integers one-to-N, but print “Fizz” if an integer is divisible by three, “Buzz” if an integer is divisible by five, and “FizzBuzz” if an integer is divisible by both three and five.

	for n := 1; n <= 100; n++ {
		if n%3 == 0 && n%5 == 0 {
			fmt.Println("FizzBuzz")
			continue
		}
		if n%3 == 0 {
			fmt.Println("Fizz")
			continue
		}
		if n%5 == 0 {
			fmt.Println("Buzz")
			continue
		}
		fmt.Println(n)
	}

	/* 4) for-range

	 */
}

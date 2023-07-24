package main

import "fmt"

/*
There is no do-while loop in Go. To emulate the C/Java code
do {
	work();
} while (condition);

you may use a for loop in one of these ways:

1) skip the init and post statemens
next := true
for next = condition {
	work()
}

2) To ensure that it executes at least once, the <condition> must be true for the first iteration and be dependent on the loop stop condition in subsequent iterations. For this we use a BOOL variable:

classic for loop:
for <initialization>; <condition>; <post-condition>
do-while loop
for next := true; next; next=<condition> {
	<loopBody>
}

3) There is also a simpler form of the do-while loop in Go, in which the <condition> of the for loop is checked after the <loopBody> and if it is not true the loop is exited:

for {
	<loopBody>
	if !<condition> { // if condition is not true
		break
	}
}
*/

func main() {
	/*
		The first version of the do-while loop:
		Print the i variable if the value is less than or equal to 5:
	*/
	i := 0
	for next := true; next; next = i <= 5 {
		fmt.Println(i)
		i++
	}

	/*
		Check that the loop is executed at least once:
	*/
	i = 10	// i is greater than 5, so the loop must finish. But the first time the bool var next is true.
	for next := true; next; next = i <= 5 {
		fmt.Println(i)
		i++
	}

	/*
		The second version of the do-while loop:
		Print the n variable if the value is less than or equal to 5:
	*/
	n := 0
	for {
		fmt.Println(n)
		n++
		if !(n <= 5) {
			break
		}
	}

	//This version of the loop is also executed at least once 😉:
	n = 10
	for {
		fmt.Println(n)
		n++
		if !(n <= 5) {
			break
		}
	}

/*
next := true
for <initialization>; <condition>; <post-condition>
for next = condition {
	work()
}
*/
n = 0
for n <= 5 {
	n++
	fmt.Println(n)
}
fmt.Println("The sum is ", n)

}

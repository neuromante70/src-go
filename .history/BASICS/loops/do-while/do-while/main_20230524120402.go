package main

import "fmt"

/*
There is no do-while loop in Go. To emulate the C/Java code
do {
	work();
} while (condition);

you may use a for loop in one of these ways:

1) skip the init and post statemens
ok := true
for ok = condition {
	work()
}

2) To ensure that it executes at least once, the <condition> must be true for the first iteration and be dependent on the loop stop condition in subsequent iterations. For this we use a BOOL variable:

for next := true; next; next=<condition> {
	<loopBody>
}

3) There is also a simpler form of the do-while loop in Go, in which the <condition> of the for loop is checked after the <loopBody> and if it is not true the loop is exited:

for {
	<loopBody>
	if !<condition> {
		break
	}
}

*/

func main() {
	/*
	1) The first version of the do-while loop:
	Print the i variable if the value is less than or equal to 5:
	*/
	i := 0
	for next := true; next; next = i <= 5 {
		fmt.Println(i)
		i++
	}

/*

*/



}

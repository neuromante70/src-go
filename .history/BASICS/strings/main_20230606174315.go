package main

import "fmt"

/*
The classic for has the form of:
for <initialization>; <condition>; <post-condition> {
	<loopBody>
}
*/

func main() {
	s := "This is a test!"
	fmt.Println("s Before the modification: ", s)
	j := "figa"
	fmt.Println(j != s)
}

package main

import "fmt"

/*
The classic for has the form of:
for <initialization>; <condition>; <post-condition> {
	<loopBody>
}
*/

func main() {
	s := "This is a string"
	fmt.Println("s Before the modification: ", s)
	j := "this modifies "
	fmt.Println(j != s)
}
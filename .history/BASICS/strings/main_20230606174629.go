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
	j := "this modifies the original string"
	fmt.Println(j != s)
	arr := [3]int{1, 2, 3}
	count := len(arr)
	for i := 0; i < count; i++ {
		arr[]
	}
}
package main

import "fmt"

/*
The classic for has the form of:
for <initialization>; <condition>; <post-condition> {
	<loopBody>
}
*/

func main() {
	s := "Te lo metto nel culo!"
	fmt.Println("s Before the modification: ", s)
	var pStr *string = &s
	fmt.Println(*pStr)
	pStr
}

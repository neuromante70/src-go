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
	fmt.Println("Before the modification: ", s)
	s = "No, sono io a mettertelo nel culo caro mio..."
	fmt.Println("After the modification: ", s)
	for i, v = range s {
		fmt.Println(s[])
	}
	s[0] = s[1]


}

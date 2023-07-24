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
	s = "No, sono io a mettertelo nel culo caro mio..."
	fmt.Println("s After the modification: ", s)

	s2 := s
	fmt.Println("s2 Before the modification: ", s2)
	s2 = "w la figa!"
	fmt.Println("s2 After the modification: ", s2, s)
	fmt.Println(&s, &s2)



}

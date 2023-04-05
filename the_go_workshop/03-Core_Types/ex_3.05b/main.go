package main

// TEXT: STRING

import "fmt"

func main() {
	comment1 := `This is the BEST
	thing ever!`
	comment2 := `This is the BEST\nthing ever!`
	comment3 := "This is the BEST thing ever!"

	fmt.Print(comment1, "\n\n")
	fmt.Print(comment2, "\n\n")
	fmt.Print(comment3, "\n")

	/* OUTPUT:
	This is the BEST
	thing ever!

	This is the BEST\nthing ever!

	This is the BEST thing ever!

	*/

	comment4 := `In "Windows" the user directory is "C:\Users\"`
	fmt.Println(comment4)
	comment5 := "In \"Windows\" the user directory is \"C:\\Users\\\""
	fmt.Println(comment5)

	/* OUTPUT:
	   In "Windows" the user directory is "C:\Users\"
	   In "Windows" the user directory is "C:\Users\"
	*/
}

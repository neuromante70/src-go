package main

import "fmt"

func main() {
	// fmt.Print("\033[2J") //clear the screen befor printing the output in the terminal
e	s := "Hello, World!"
	fmt.Printf("%s\n", s)
	fmt.Printf("%q\n", s)
}

package main

import "fmt"

func main() {
	fmt.Print("\033[2J") //clear the screen befor printing the output in the terminal
	fmt.Printf("hello, world\n")
}

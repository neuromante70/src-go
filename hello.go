package main

import "fmt"

func function(int) int {
	return 0
}

func main() {
	// fmt.Print("\033[2J") //clear the screen befor printing the output in the terminal
	result := function(1)
	fmt.Println("\nYour function result is:\n", result)
}

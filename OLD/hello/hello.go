package main

import (
"fmt"
"errors"
)

func function(int) int {
	return 0
}

func main() {
	fmt.Print("\033[2J") //clear the screen befor printing the output in the terminal
	fmt.Println("In culo alla balena")
  result := function(1)
	fmt.Println("\nYour function result is:\n", result)
}

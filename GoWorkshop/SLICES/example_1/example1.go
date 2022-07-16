package main

import "fmt"

func function(int) int {
	return 0
}

func inCuloAtuaMadre(int) {
	return
}

func relayArray(int) int {
	return 0
}

func main() {
	// fmt.Print("\033[2J") //clear the screen befor printing the output in the terminal
	relayArray(5)
	myVar := function(5)
	var x = [...]int{1, 2, 3}
	var y = [3]int{1, 2, 3}
	fmt.Println(x == y)
	fmt.Println("Capacity of x:", cap(x), "Capacity of y:", cap(y))
	fmt.Println("Myvar is: ", myVar)
}

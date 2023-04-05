package main

import "fmt"


func main() {
	// fmt.Print("\033[2J") //clear the screen befor printing the output in the terminal
	var x = [...]int{1, 2, 3}
	var y = [3]int{1, 2, 3}
	fmt.Println(x == y)
	fmt.Println("Capacity of x:", cap(x), "Capacity of y:", cap(y))
}

package main

import "fmt"

func main() {
	x := make([]int, 5, 10)
	x = append(x, 10)
	fmt.Println(x)
}

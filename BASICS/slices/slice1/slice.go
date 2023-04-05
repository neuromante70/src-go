package main

import "fmt"

func main() {
	var x = []int{1, 2, 3}
	// fmt.Println(x == nil)
	y := []int{20, 30, 40}
	x = append(x, y...)
	fmt.Println(x)
	fmt.Printf("%#v %T", x, x)

}

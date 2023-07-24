package main

import (
	"fmt"
	"strconv"
)

func main() {
	var title string
	var copies int
	title = "For the Love of Go"
	copies = 99
	fmt.Println(title)
	fmt.Println(copies)
	title = strconv.Itoa(copies)
	fmt.Printf("this is a %T\n", title)
}

package main

import "fmt"

func foo() {
	fmt.Println("test with ")
}

func main() {
	var title string
	var copies int
	title = "For the Love of go"
	copies = 99
	foo()
	fmt.Println("vim-go")
	
}

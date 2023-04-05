package main

import "fmt"

func foo() {
	fmt.Println("test with ")
}

func main() {
	var title, author string
	var copies int
	title = "For the Love of go"
	author = "John Arundel"
	copies = title
	foo()
	fmt.Println(author, "-", title)
	fmt.Println(copies)
}

package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Copies int
}

func GetBook(catalog map[int]Book, ID int) Book {
	return catalog[ID]
	}

func foo() {
	fmt.Println("test with ")
}

func main() {
	foo()
	fmt.Println("vim-go")
}

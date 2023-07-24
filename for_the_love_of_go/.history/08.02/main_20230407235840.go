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


func main() {
	catalog := map[int]Book{
		1: {ID: 1, Title: "For the Love of Go"},
		2: {ID: 2, Title: "The Power of Go: Tools"},
		}
	fmt.Println("vim-go")
}

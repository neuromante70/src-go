package main

import "fmt"

type Book struct {
	Title  string
	Author string
	Copies int
}

func main() {
	fmt.Println("vim-go")
	var books = []Book
	// books = []Book{}
	var books = []Book{
		{Title: "Delightfully Uneventful Trip on the Orient Express"},
		{Title: "One Hundred Years of Good Company"},
	}
	first := books[0]
	fmt.Println(len(books))
	fmt.Println(first)

}

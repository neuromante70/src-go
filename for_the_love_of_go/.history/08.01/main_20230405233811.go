package main

import (
	"fmt"
)

type Book struct {
	Title  string
	Author string
	Copies int
}

func main() {
	// var books []Book
	// books = []Book{}
	books := []Book{
		{Title: "Delightfully Uneventful Trip on the Orient Express"},
		{Title: "One Hundred Years of Good Company"},
	}
	first := books[0]
	// first = Book{Title: "Heart of Kindness"}
	books[0].Title = "Heart of Kindness"
	fmt.Println(len(books))
	b := Book{Title: "The Grapes of Mild Irritation"}
	books = append(books, b)
	fmt.Println(books)
}

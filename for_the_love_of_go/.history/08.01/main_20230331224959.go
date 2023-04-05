package main

import "fmt"

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
	first = Book
	fmt.Println(len(books))
	fmt.Println(first)

}

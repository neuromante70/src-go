package main

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
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
	fmt.Println(first)

	want := []string{"same", "same", "same"}
	got := []
}
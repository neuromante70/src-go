package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	var books []Book
	books = []Books{
		{Title:  "Delightfully Uneventful Trip on the Orient Express"},
		{Title: "One Hundred Years of Good Company"},
	}
	first := books[0]
	fmt.Println(len(books))


}

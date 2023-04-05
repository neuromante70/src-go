package bookstore

import (
	"fmt"
)

// Book represents information about a book.
type Book struct {
	Title  string
	Author string
	Copies int
}

b := bookstore.Book{
	Title: "Nicholas Chuckleby",
	Author: "Charles Dickens",
	Copies: 8,
}

fmt.Println(b)
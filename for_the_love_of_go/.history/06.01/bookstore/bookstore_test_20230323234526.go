package bookstore_test

import (
	"bookstore"
	"testing"
)

// bookstore.Book{
// 	Title:	"Nicholas Chuckleby",
// 	Author:	"Charles Dickens",
// 	Copies:	8
// }

func TestBook(t *testing.T) {
	t.Parallel()

	_ = bookstore.Book{
		Title:  "Spak Joy",
		Author: "Marie Kondo",
		Copies: 2,
	}
}

func TestBucioDiCulo(t ) {
	dimension	string
	number		int
	clean		bool
}
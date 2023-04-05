package bookstore_test

import (
	"bookstore"
	"testing"
	"the_go_workshop/Go_books/For_the_Love_of_Go/for_the_love_of_go-code/06.01/bookstore"
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

func TestBucioDiCulo(t *testing.T) {
	t.Parallel()
	_ = bookstore.BucioDiCulo{
	dimension	string
	number		int
	clean		bool
}
package bookstore_test

import (
	"bookstore"
	"for_the_love_of_go/06.01/bookstore"
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

func TestBuy(t *testing.T) {
	t.Parallel()

	b := bookstore.Book{
		Title:		
	}
}
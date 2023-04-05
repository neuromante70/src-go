package bookstore_test

import (
	"bookstore"
	"testing"
)

func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondō",
		Copies: 2,
	}
}

func TestBuy(t *testing.T) {
	b := bookstore.Book{
		Title: "Nicholas Chuckleby",
		Author: "Charles Dickens",
		Copies: 8,
	}
	
}